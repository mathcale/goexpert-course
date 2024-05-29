package handlers

import (
	"encoding/json"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"github.com/mathcale/goexpert-course/otel-lab/internal/entities/dto"
	"github.com/mathcale/goexpert-course/otel-lab/internal/pkg/responsehandler"
	"github.com/mathcale/goexpert-course/otel-lab/internal/usecases/input"
)

type WebInputHandlerInterface interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type WebInputHandler struct {
	ResponseHandler responsehandler.WebResponseHandlerInterface
	InputUseCase    input.InputUseCaseInterface
	Tracer          trace.Tracer
}

func NewWebInputHandler(
	rh responsehandler.WebResponseHandlerInterface,
	inputUC input.InputUseCaseInterface,
	tracer trace.Tracer,
) *WebInputHandler {
	return &WebInputHandler{
		ResponseHandler: rh,
		InputUseCase:    inputUC,
		Tracer:          tracer,
	}
}

func (h *WebInputHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var dto dto.InputUCInput

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)
	ctx, span := h.Tracer.Start(ctx, "climate")
	defer span.End()

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		span.SetStatus(codes.Error, "invalid input")
		span.RecordError(err)
		span.End()

		h.ResponseHandler.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(r.Header))

	input, err := h.InputUseCase.Execute(ctx, dto)
	if err != nil {
		span.SetStatus(codes.Error, "internal error")
		span.RecordError(err)
		span.End()

		h.ResponseHandler.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	h.ResponseHandler.Respond(w, http.StatusOK, input)
}

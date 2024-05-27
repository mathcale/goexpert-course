package handlers

import (
	"encoding/json"
	"net/http"

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
}

func NewWebInputHandler(
	rh responsehandler.WebResponseHandlerInterface,
	inputUC input.InputUseCaseInterface,
) *WebInputHandler {
	return &WebInputHandler{
		ResponseHandler: rh,
		InputUseCase:    inputUC,
	}
}

func (h *WebInputHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var dto dto.InputUCInput

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		h.ResponseHandler.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	input, err := h.InputUseCase.Execute(dto)
	if err != nil {
		h.ResponseHandler.RespondWithError(w, http.StatusInternalServerError, err)
	}

	h.ResponseHandler.Respond(w, http.StatusOK, input)
}

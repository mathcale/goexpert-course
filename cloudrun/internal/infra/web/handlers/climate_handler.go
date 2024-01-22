package handlers

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/mathcale/goexpert-course/cloudrun/internal/entities/dto"
	"github.com/mathcale/goexpert-course/cloudrun/internal/pkg/responsehandler"
)

type WebClimateHandler struct {
	ResponseHandler responsehandler.WebResponseHandlerInterface
}

func NewWebClimateHandler(rh responsehandler.WebResponseHandlerInterface) *WebClimateHandler {
	return &WebClimateHandler{
		ResponseHandler: rh,
	}
}

func (h *WebClimateHandler) GetTemperaturesByZipCode(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	zipStr := qs.Get("zipcode")

	if zipStr == "" {
		h.ResponseHandler.RespondWithError(w, http.StatusUnprocessableEntity, errors.New("invalid zipcode"))
		return
	}

	matched, err := regexp.MatchString(`\d{5}[\-]?\d{3}`, zipStr)

	if !matched || err != nil {
		h.ResponseHandler.RespondWithError(w, http.StatusUnprocessableEntity, errors.New("invalid zipcode"))
		return
	}

	// TODO: call use-cases here

	h.ResponseHandler.Respond(w, http.StatusOK, dto.GetTemperaturesByZipCodeOutput{
		Celcius:    0,
		Fahrenheit: 0,
		Kelvin:     0,
	})
}

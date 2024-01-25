package handlers

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/mathcale/goexpert-course/cloudrun/internal/entities/dto"
	"github.com/mathcale/goexpert-course/cloudrun/internal/pkg/responsehandler"
	"github.com/mathcale/goexpert-course/cloudrun/internal/usecases/climate"
	"github.com/mathcale/goexpert-course/cloudrun/internal/usecases/location"
)

type WebClimateHandlerInterface interface {
	GetTemperaturesByZipCode(w http.ResponseWriter, r *http.Request)
}

type WebClimateHandler struct {
	ResponseHandler              responsehandler.WebResponseHandlerInterface
	FindLocationByZipCodeUseCase location.FindByZipCodeUseCaseInterface
	FindClimateByCityNameUseCase climate.FindByCityNameUseCaseInterface
}

func NewWebClimateHandler(
	rh responsehandler.WebResponseHandlerInterface,
	findByZipCodeUC location.FindByZipCodeUseCaseInterface,
	findByCityNameUC climate.FindByCityNameUseCaseInterface,
) *WebClimateHandler {
	return &WebClimateHandler{
		ResponseHandler:              rh,
		FindLocationByZipCodeUseCase: findByZipCodeUC,
		FindClimateByCityNameUseCase: findByCityNameUC,
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

	location, err := h.FindLocationByZipCodeUseCase.Execute(zipStr)
	if err != nil {
		h.ResponseHandler.RespondWithError(w, http.StatusInternalServerError, err)
	}

	if location.City == "" {
		h.ResponseHandler.RespondWithError(w, http.StatusNotFound, errors.New("zipcode not found"))
		return
	}

	climate, err := h.FindClimateByCityNameUseCase.Execute(location.City)
	if err != nil {
		h.ResponseHandler.RespondWithError(w, http.StatusInternalServerError, err)
	}

	fahrenheit := climate.Current.TempC*1.8 + 32
	kelvin := climate.Current.TempC + 273.15

	h.ResponseHandler.Respond(w, http.StatusOK, dto.GetTemperaturesByZipCodeOutput{
		Celcius:    float32(climate.Current.TempC),
		Fahrenheit: float32(fahrenheit),
		Kelvin:     float32(kelvin),
	})
}

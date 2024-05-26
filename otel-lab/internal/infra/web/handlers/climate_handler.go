package handlers

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/mathcale/goexpert-course/otel-lab/internal/entities/dto"
	"github.com/mathcale/goexpert-course/otel-lab/internal/pkg/responsehandler"
	"github.com/mathcale/goexpert-course/otel-lab/internal/usecases/climate"
	"github.com/mathcale/goexpert-course/otel-lab/internal/usecases/location"
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

	if err := validateInput(zipStr); err != nil {
		h.ResponseHandler.RespondWithError(w, http.StatusUnprocessableEntity, err)
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

	fahrenheit, kelvin := convertTemperature(climate.Current.TempC)

	h.ResponseHandler.Respond(w, http.StatusOK, dto.GetTemperaturesByZipCodeOutput{
		Celcius:    float32(climate.Current.TempC),
		Fahrenheit: float32(fahrenheit),
		Kelvin:     float32(kelvin),
	})
}

func validateInput(zipcode string) error {
	if zipcode == "" {
		return errors.New("invalid zipcode")
	}

	matched, err := regexp.MatchString(`\b\d{5}[\-]?\d{3}\b`, zipcode)
	if !matched || err != nil {
		return errors.New("invalid zipcode")
	}

	return nil
}

func convertTemperature(celcius float64) (float64, float64) {
	fahrenheit := celcius*1.8 + 32
	kelvin := celcius + 273.15

	return fahrenheit, kelvin
}

package location

import (
	"fmt"

	"github.com/rs/zerolog"

	"github.com/mathcale/goexpert-course/otel-lab/internal/entities"
	"github.com/mathcale/goexpert-course/otel-lab/internal/pkg/httpclient"
)

type FindByZipCodeUseCaseInterface interface {
	Execute(zipCode string) (*entities.Location, error)
}

type FindByZipCodeUseCase struct {
	HttpClient httpclient.HttpClientInterface
	Logger     zerolog.Logger
}

func NewFindByZipCodeUseCase(
	httpClient httpclient.HttpClientInterface,
	logger zerolog.Logger,
) *FindByZipCodeUseCase {
	return &FindByZipCodeUseCase{
		HttpClient: httpClient,
		Logger:     logger,
	}
}

func (uc *FindByZipCodeUseCase) Execute(zipCode string) (*entities.Location, error) {
	var location entities.Location

	uc.Logger.Info().Msgf("[FindByZipCode] Calling API with zipcode [%s]", zipCode)

	if err := uc.HttpClient.Get(fmt.Sprintf("/%s/json/", zipCode), &location); err != nil {
		return nil, err.Error
	}

	uc.Logger.Debug().Msgf("[FindByZipCode] Got location [%+v]", location)

	return &location, nil
}

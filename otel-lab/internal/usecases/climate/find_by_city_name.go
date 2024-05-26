package climate

import (
	"fmt"
	"net/url"

	"github.com/rs/zerolog"

	"github.com/mathcale/goexpert-course/otel-lab/internal/entities"
	"github.com/mathcale/goexpert-course/otel-lab/internal/pkg/httpclient"
)

type FindByCityNameUseCaseInterface interface {
	Execute(city string) (*entities.Climate, error)
}

type FindByCityNameUseCase struct {
	HttpClient httpclient.HttpClientInterface
	Logger     zerolog.Logger
	APIKey     string
}

func NewFindByCityNameUseCase(
	httpClient httpclient.HttpClientInterface,
	logger zerolog.Logger,
	apiKey string,
) *FindByCityNameUseCase {
	return &FindByCityNameUseCase{
		HttpClient: httpClient,
		Logger:     logger,
		APIKey:     apiKey,
	}
}

func (uc *FindByCityNameUseCase) Execute(city string) (*entities.Climate, error) {
	var climate entities.Climate

	uc.Logger.Info().Msgf("[FindByCityName] Calling API with city name [%s]", city)

	if err := uc.HttpClient.Get(fmt.Sprintf("/v1/current.json?key=%s&q=%s&aqi=no", uc.APIKey, url.QueryEscape(city)), &climate); err != nil {
		return nil, err.Error
	}

	uc.Logger.Debug().Msgf("[FindByCityName] Got climate data [%+v]", climate)

	return &climate, nil
}

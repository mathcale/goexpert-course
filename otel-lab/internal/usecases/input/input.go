package input

import (
	"fmt"

	"github.com/rs/zerolog"

	"github.com/mathcale/goexpert-course/otel-lab/internal/entities/dto"
	"github.com/mathcale/goexpert-course/otel-lab/internal/pkg/httpclient"
)

type InputUseCaseInterface interface {
	Execute(input dto.InputUCInput) (*dto.GetTemperaturesByZipCodeOutput, error)
}

type InputUseCase struct {
	HttpClient httpclient.HttpClientInterface
	Logger     zerolog.Logger
}

func NewInputUseCase(
	httpClient httpclient.HttpClientInterface,
	logger zerolog.Logger,
) *InputUseCase {
	return &InputUseCase{
		HttpClient: httpClient,
		Logger:     logger,
	}
}

func (uc *InputUseCase) Execute(input dto.InputUCInput) (*dto.GetTemperaturesByZipCodeOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	uc.Logger.Info().Msgf("[Input] Calling Orchestrator API with zipcode [%s]", input.Zipcode)

	var response dto.GetTemperaturesByZipCodeOutput

	if err := uc.HttpClient.Get(fmt.Sprintf("/?zipcode=%s", input.Zipcode), &response); err != nil {
		return nil, err.Error
	}

	uc.Logger.Debug().Msgf("[Input] Got data: %+v", response)

	return &response, nil
}

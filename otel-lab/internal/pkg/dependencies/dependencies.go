package dependencies

import (
	"time"

	"github.com/mathcale/goexpert-course/otel-lab/config"
	"github.com/mathcale/goexpert-course/otel-lab/internal/infra/web"
	"github.com/mathcale/goexpert-course/otel-lab/internal/infra/web/handlers"
	"github.com/mathcale/goexpert-course/otel-lab/internal/pkg/httpclient"
	"github.com/mathcale/goexpert-course/otel-lab/internal/pkg/logger"
	"github.com/mathcale/goexpert-course/otel-lab/internal/pkg/responsehandler"
	"github.com/mathcale/goexpert-course/otel-lab/internal/usecases/climate"
	"github.com/mathcale/goexpert-course/otel-lab/internal/usecases/input"
	"github.com/mathcale/goexpert-course/otel-lab/internal/usecases/location"
)

type InputServiceDependencies struct {
	WebServer web.WebServerInterface
}

type OrchestratorServiceDependencies struct {
	WebServer web.WebServerInterface
}

type sharedDependencies struct {
	ResponseHandler   responsehandler.WebResponseHandler
	Logger            logger.Logger
	HttpClientTimeout time.Duration
}

func ResolveInputServiceDependencies(config *config.Conf) InputServiceDependencies {
	sharedDeps := resolveSharedDependencies(config)

	httpClient := httpclient.NewHttpClient(config.OrchestratorServiceHost, sharedDeps.HttpClientTimeout)

	inputUC := input.NewInputUseCase(httpClient, sharedDeps.Logger.GetLogger())

	webInputHandler := handlers.NewWebInputHandler(&sharedDeps.ResponseHandler, inputUC)

	webRouter := web.NewInputWebRouter(webInputHandler)
	webServer := web.NewWebServer(config.InputServiceWebServerPort, sharedDeps.Logger.GetLogger(), webRouter.Build())

	return InputServiceDependencies{
		WebServer: webServer,
	}
}

func ResolveOrchestratorServiceDependencies(config *config.Conf) OrchestratorServiceDependencies {
	sharedDeps := resolveSharedDependencies(config)

	viaCepAPIHttpClient := httpclient.NewHttpClient(config.ViaCepApiBaseUrl, sharedDeps.HttpClientTimeout)
	weatherAPIHttpClient := httpclient.NewHttpClient(config.WeatherApiBaseUrl, sharedDeps.HttpClientTimeout)

	findByZipCodeUseCase := location.NewFindByZipCodeUseCase(viaCepAPIHttpClient, sharedDeps.Logger.GetLogger())
	findByCityNameUseCase := climate.NewFindByCityNameUseCase(weatherAPIHttpClient, sharedDeps.Logger.GetLogger(), config.WeatherApiKey)

	webClimateHandler := handlers.NewWebClimateHandler(&sharedDeps.ResponseHandler, findByZipCodeUseCase, findByCityNameUseCase)

	webRouter := web.NewOrchestratorWebRouter(webClimateHandler)
	webServer := web.NewWebServer(config.OrchestratorServiceWebServerPort, sharedDeps.Logger.GetLogger(), webRouter.Build())

	return OrchestratorServiceDependencies{
		WebServer: webServer,
	}
}

func resolveSharedDependencies(config *config.Conf) sharedDependencies {
	logger := logger.NewLogger(config.LogLevel)
	logger.Setup()

	responseHandler := responsehandler.NewWebResponseHandler()

	httpClientTimeout := time.Duration(config.HttpClientTimeout) * time.Millisecond

	return sharedDependencies{
		ResponseHandler:   *responseHandler,
		Logger:            *logger,
		HttpClientTimeout: httpClientTimeout,
	}
}

package dependencies

import (
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

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
	ServiceName string
	WebServer   web.WebServerInterface
}

type OrchestratorServiceDependencies struct {
	ServiceName string
	WebServer   web.WebServerInterface
}

type sharedDependencies struct {
	ResponseHandler   responsehandler.WebResponseHandler
	Logger            logger.Logger
	HttpClientTimeout time.Duration
	Tracer            trace.Tracer
}

func ResolveInputServiceDependencies(config *config.Conf) InputServiceDependencies {
	serviceName := "input-service"
	sharedDeps := resolveSharedDependencies(config, serviceName)

	httpClient := httpclient.NewHttpClient(config.OrchestratorServiceHost, sharedDeps.HttpClientTimeout)

	inputUC := input.NewInputUseCase(httpClient, sharedDeps.Logger.GetLogger())

	webInputHandler := handlers.NewWebInputHandler(&sharedDeps.ResponseHandler, inputUC, sharedDeps.Tracer)

	webRouter := web.NewInputWebRouter(webInputHandler)
	webServer := web.NewWebServer(config.InputServiceWebServerPort, sharedDeps.Logger.GetLogger(), webRouter.Build())

	return InputServiceDependencies{
		ServiceName: serviceName,
		WebServer:   webServer,
	}
}

func ResolveOrchestratorServiceDependencies(config *config.Conf) OrchestratorServiceDependencies {
	serviceName := "orchestrator-service"
	sharedDeps := resolveSharedDependencies(config, serviceName)

	viaCepAPIHttpClient := httpclient.NewHttpClient(config.ViaCepApiBaseUrl, sharedDeps.HttpClientTimeout)
	weatherAPIHttpClient := httpclient.NewHttpClient(config.WeatherApiBaseUrl, sharedDeps.HttpClientTimeout)

	findByZipCodeUseCase := location.NewFindByZipCodeUseCase(viaCepAPIHttpClient, sharedDeps.Logger.GetLogger())
	findByCityNameUseCase := climate.NewFindByCityNameUseCase(weatherAPIHttpClient, sharedDeps.Logger.GetLogger(), config.WeatherApiKey)

	webClimateHandler := handlers.NewWebClimateHandler(&sharedDeps.ResponseHandler, findByZipCodeUseCase, findByCityNameUseCase, sharedDeps.Tracer)

	webRouter := web.NewOrchestratorWebRouter(webClimateHandler)
	webServer := web.NewWebServer(config.OrchestratorServiceWebServerPort, sharedDeps.Logger.GetLogger(), webRouter.Build())

	return OrchestratorServiceDependencies{
		ServiceName: serviceName,
		WebServer:   webServer,
	}
}

func resolveSharedDependencies(config *config.Conf, serviceName string) sharedDependencies {
	logger := logger.NewLogger(config.LogLevel)
	logger.Setup()

	responseHandler := responsehandler.NewWebResponseHandler()

	httpClientTimeout := time.Duration(config.HttpClientTimeout) * time.Millisecond

	tracer := otel.Tracer(serviceName)

	return sharedDependencies{
		ResponseHandler:   *responseHandler,
		Logger:            *logger,
		HttpClientTimeout: httpClientTimeout,
		Tracer:            tracer,
	}
}

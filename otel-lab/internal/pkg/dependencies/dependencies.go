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
	"github.com/mathcale/goexpert-course/otel-lab/internal/usecases/location"
)

type OrchestratorServiceDependencies struct {
	WebServer web.WebServerInterface
}

func ResolveOrchestratorServiceDependencies(config *config.Conf) OrchestratorServiceDependencies {
	responseHandler := responsehandler.NewWebResponseHandler()

	logger := logger.NewLogger(config.LogLevel)
	logger.Setup()

	httpClientTimeout := time.Duration(config.HttpClientTimeout) * time.Millisecond
	viaCepAPIHttpClient := httpclient.NewHttpClient(config.ViaCepApiBaseUrl, httpClientTimeout)
	weatherAPIHttpClient := httpclient.NewHttpClient(config.WeatherApiBaseUrl, httpClientTimeout)

	findByZipCodeUseCase := location.NewFindByZipCodeUseCase(viaCepAPIHttpClient, logger.GetLogger())
	findByCityNameUseCase := climate.NewFindByCityNameUseCase(weatherAPIHttpClient, logger.GetLogger(), config.WeatherApiKey)

	webClimateHandler := handlers.NewWebClimateHandler(responseHandler, findByZipCodeUseCase, findByCityNameUseCase)

	webRouter := web.NewWebRouter(webClimateHandler)
	webServer := web.NewWebServer(config.OrchestratorServiceWebServerPort, logger.GetLogger(), webRouter.Build())

	return OrchestratorServiceDependencies{
		WebServer: webServer,
	}
}

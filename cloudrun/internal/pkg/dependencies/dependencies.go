package dependencies

import (
	"time"

	"github.com/mathcale/goexpert-course/cloudrun/config"
	"github.com/mathcale/goexpert-course/cloudrun/internal/infra/web"
	"github.com/mathcale/goexpert-course/cloudrun/internal/infra/web/handlers"
	"github.com/mathcale/goexpert-course/cloudrun/internal/pkg/httpclient"
	"github.com/mathcale/goexpert-course/cloudrun/internal/pkg/logger"
	"github.com/mathcale/goexpert-course/cloudrun/internal/pkg/responsehandler"
	"github.com/mathcale/goexpert-course/cloudrun/internal/usecases/climate"
	"github.com/mathcale/goexpert-course/cloudrun/internal/usecases/location"
)

type AppDependencies struct {
	Logger           logger.LoggerInterface
	ResponseHandler  responsehandler.WebResponseHandlerInterface
	ViaCepHttpClient httpclient.HttpClientInterface
	WebServer        web.WebServerInterface

	// Handlers
	WebClimateHandler handlers.WebClimateHandlerInterface

	// Use-cases
	FindByZipCodeUseCase  location.FindByZipCodeUseCaseInterface
	FindByCityNameUseCase climate.FindByCityNameUseCaseInterface
}

func Build(config *config.Conf) AppDependencies {
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
	webServer := web.NewWebServer(config.WebServerPort, logger.GetLogger(), webRouter.Build())

	return AppDependencies{
		Logger:           logger,
		ResponseHandler:  responseHandler,
		ViaCepHttpClient: viaCepAPIHttpClient,
		WebServer:        webServer,

		WebClimateHandler: webClimateHandler,

		FindByZipCodeUseCase:  findByZipCodeUseCase,
		FindByCityNameUseCase: findByCityNameUseCase,
	}
}

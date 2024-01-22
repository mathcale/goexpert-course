package main

import (
	"github.com/mathcale/goexpert-course/cloudrun/config"
	"github.com/mathcale/goexpert-course/cloudrun/internal/infra/web"
)

func main() {
	configs, configsErr := config.LoadConfig(".")
	if configsErr != nil {
		panic(configsErr)
	}

	logger := NewLogger(configs.LogLevel)
	logger.Setup()

	webServer := NewWebServer(configs.WebServerPort, logger.GetLogger(), createWebHandlers())
	webServer.Start()
}

func createWebHandlers() []web.RouteHandler {
	webClimateHandler := NewWebClimateHandler()

	return []web.RouteHandler{
		{
			Path:        "/",
			Method:      "GET",
			HandlerFunc: webClimateHandler.GetTemperaturesByZipCode,
		},
	}
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mathcale/goexpert-course/cloudrun/internal/infra/web"
	"github.com/mathcale/goexpert-course/cloudrun/internal/infra/web/handlers"
	"github.com/mathcale/goexpert-course/cloudrun/internal/pkg/logger"
	"github.com/mathcale/goexpert-course/cloudrun/internal/pkg/responsehandler"
	"github.com/rs/zerolog"
)

// Injectors from wire.go:

func NewLogger(level string) *logger.Logger {
	loggerLogger := logger.NewLogger(level)
	return loggerLogger
}

func NewWebClimateHandler() *handlers.WebClimateHandler {
	webResponseHandler := responsehandler.NewWebResponseHandler()
	webClimateHandler := handlers.NewWebClimateHandler(webResponseHandler)
	return webClimateHandler
}

func NewWebServer(port string, logger2 zerolog.Logger, handlers2 []web.RouteHandler) *web.WebServer {
	webServer := web.NewWebServer(port, logger2, handlers2)
	return webServer
}

// wire.go:

var setResponseHandlerDependency = wire.NewSet(responsehandler.NewWebResponseHandler, wire.Bind(
	new(responsehandler.WebResponseHandlerInterface),
	new(*responsehandler.WebResponseHandler),
),
)

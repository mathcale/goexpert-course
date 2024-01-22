//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rs/zerolog"

	"github.com/mathcale/goexpert-course/cloudrun/internal/infra/web"
	"github.com/mathcale/goexpert-course/cloudrun/internal/infra/web/handlers"
	"github.com/mathcale/goexpert-course/cloudrun/internal/pkg/logger"
	"github.com/mathcale/goexpert-course/cloudrun/internal/pkg/responsehandler"
)

var setResponseHandlerDependency = wire.NewSet(
	responsehandler.NewWebResponseHandler,
	wire.Bind(
		new(responsehandler.WebResponseHandlerInterface),
		new(*responsehandler.WebResponseHandler),
	),
)

func NewLogger(level string) *logger.Logger {
	wire.Build(logger.NewLogger)
	return &logger.Logger{}
}

func NewWebClimateHandler() *handlers.WebClimateHandler {
	wire.Build(
		setResponseHandlerDependency,
		handlers.NewWebClimateHandler,
	)

	return &handlers.WebClimateHandler{}
}

func NewWebServer(port string, logger zerolog.Logger, handlers []web.RouteHandler) *web.WebServer {
	wire.Build(web.NewWebServer)

	return &web.WebServer{}
}

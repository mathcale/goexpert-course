package web

import (
	"net/http"

	"github.com/mathcale/goexpert-course/otel-lab/internal/infra/web/handlers"
)

type WebRouterInterface interface {
	Build() []RouteHandler
}

type WebRouter struct {
	WebClimateHandler handlers.WebClimateHandlerInterface
}

func NewWebRouter(webClimateHandler handlers.WebClimateHandlerInterface) *WebRouter {
	return &WebRouter{
		WebClimateHandler: webClimateHandler,
	}
}

func (wr *WebRouter) Build() []RouteHandler {
	return []RouteHandler{
		{
			Path:        "/",
			Method:      http.MethodGet,
			HandlerFunc: wr.WebClimateHandler.GetTemperaturesByZipCode,
		},
	}
}

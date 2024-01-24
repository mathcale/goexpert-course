package web

import "github.com/mathcale/goexpert-course/cloudrun/internal/infra/web/handlers"

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
			Method:      "GET",
			HandlerFunc: wr.WebClimateHandler.GetTemperaturesByZipCode,
		},
	}
}

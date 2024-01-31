package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	chizero "github.com/ironstar-io/chizerolog"
	"github.com/rs/zerolog"
)

type WebServerInterface interface {
	Start()
}

type RouteHandler struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

type WebServer struct {
	Router        chi.Router
	Handlers      []RouteHandler
	WebServerPort int
	Logger        zerolog.Logger
}

func NewWebServer(serverPort int, logger zerolog.Logger, handlers []RouteHandler) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      handlers,
		WebServerPort: serverPort,
		Logger:        logger,
	}
}

func (s *WebServer) Start() {
	s.Router.Use(chizero.LoggerMiddleware(&s.Logger))

	for _, h := range s.Handlers {
		s.Logger.Debug().Msgf("Registering route %s %s", h.Method, h.Path)
		s.Router.MethodFunc(h.Method, h.Path, h.HandlerFunc)
	}

	s.Logger.Info().Msgf("Starting server on port %d", s.WebServerPort)

	http.ListenAndServe(fmt.Sprintf(":%d", s.WebServerPort), s.Router)
}

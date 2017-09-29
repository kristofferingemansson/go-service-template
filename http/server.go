package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kristofferingemansson/go-service-template/pkg"
	"net/http"
)

// Server ..
type Server interface {
	Listen(addr string, static RoutableHandler, api RoutableHandler, ws RoutableHandler) error
}

type server struct {
	logger pkg.Logger
}

// RoutableHandler ..
type RoutableHandler interface {
	Route(chi.Router)
}

// NewServer ..
func NewServer(logger pkg.Logger) Server {
	return &server{
		logger: logger,
	}
}

func (s *server) Listen(addr string, static RoutableHandler, api RoutableHandler, ws RoutableHandler) error {
	router := chi.NewRouter()
	router.Use(
		middleware.Recoverer,
		middleware.Logger,
	)

	router.Route("/", static.Route)
	router.Route("/api", api.Route)
	router.Route("/ws", ws.Route)

	s.logger.Log(
		"msg", "http.server.listen",
		"addr", addr,
	)

	return http.ListenAndServe(addr, router)
}

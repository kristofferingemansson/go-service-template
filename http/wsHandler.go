package http

import (
	"github.com/go-chi/chi"
	"github.com/kristofferingemansson/go-service-template/pkg"
	"net/http"
)

type wsHandler struct {
	logger pkg.Logger
}

// NewWsHandler ..
func NewWsHandler(logger pkg.Logger) RoutableHandler {
	return &wsHandler{
		logger: logger,
	}
}

func (h *wsHandler) Route(router chi.Router) {
	router.Get("/", http.NotFound)
}

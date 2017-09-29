package http

import (
	"github.com/go-chi/chi"
	"github.com/kristofferingemansson/go-service-template/pkg"
	"net/http"
)

type apiHandler struct {
	logger pkg.Logger
}

// NewAPIHandler ..
func NewAPIHandler(logger pkg.Logger) RoutableHandler {
	return &apiHandler{
		logger: logger,
	}
}

func (h *apiHandler) Route(router chi.Router) {
	router.Get("/", http.NotFound)
}

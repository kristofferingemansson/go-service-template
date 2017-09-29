package http

import (
	"github.com/go-chi/chi"
	"github.com/kristofferingemansson/go-service-template/pkg"
	"net/http"
)

type staticHandler struct {
	logger pkg.Logger
	root   string
}

// NewStaticHandler ..
func NewStaticHandler(logger pkg.Logger, root string) RoutableHandler {
	return &staticHandler{
		logger: logger,
		root:   root,
	}
}

func (h *staticHandler) Route(router chi.Router) {
	router.Get("/", http.FileServer(http.Dir(h.root)).ServeHTTP)
}

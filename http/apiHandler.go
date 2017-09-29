package http

import (
	"github.com/go-chi/chi"
	"github.com/kristofferingemansson/go-service-template/pkg"
	"github.com/kristofferingemansson/go-service-template/quote"
	"net/http"
)

type apiHandler struct {
	logger  pkg.Logger
	service quote.Service
}

// NewAPIHandler ..
func NewAPIHandler(logger pkg.Logger, service quote.Service) RoutableHandler {
	return &apiHandler{
		logger:  logger,
		service: service,
	}
}

func (h *apiHandler) Route(router chi.Router) {
	router.Get("/quote", h.getQuote)
}

func (h *apiHandler) getQuote(w http.ResponseWriter, r *http.Request) {
	q, err := h.service.GenerateQuote()
	if err != nil {
		encodeError(w, err)
		return
	}

	encodeResponse(w, map[string]interface{}{
		"quote": q,
	})
}

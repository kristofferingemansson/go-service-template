package http

import (
	"github.com/go-chi/chi"
	"github.com/kristofferingemansson/go-service-template/pkg"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	handler := NewStaticHandler(pkg.NilLogger, "")

	router := chi.NewRouter()
	handler.Route(router)

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Error("Unexpected status code")
	}

	r = httptest.NewRequest(http.MethodGet, "/mostlikelymissingurl", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if w.Code != http.StatusNotFound {
		t.Error("Unexpected status code")
	}
}

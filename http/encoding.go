package http

import (
	"encoding/json"
	"github.com/kristofferingemansson/go-service-template/pkg"
	"net/http"
)

func encodeError(w http.ResponseWriter, err error) {
	switch err {
	case pkg.ErrNotImplemented:
		w.WriteHeader(http.StatusNotImplemented)
	default:
		w.WriteHeader(500)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func encodeResponse(w http.ResponseWriter, r interface{}) {
	if r != nil {
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(r)
	}
}

package rest

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(func(next http.Handler) http.Handler {
		return http.MaxBytesHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		}), 4096)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"alive": true,
		})
	})

	return r
}

package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func newRouteur() *chi.Mux {
	r := chi.NewRouter()

	// Middlewares essentiels
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	return r
}

func NewApi() *chi.Mux {
	r := newRouteur()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		root(w)
	})

	r.Route("/api/gestion", func(r chi.Router) {
		r.Get("/accountValue/{id}", func(w http.ResponseWriter, r *http.Request) {
			accountValue(w, r)
		})
	})

	return r
}

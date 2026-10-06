package api

import (
	"database/sql"
	"net/http"
	"time"

	middleware2 "github.com/Chefmine8/OpenCroupier/internal/middleware"
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

func NewApi(db *sql.DB) *chi.Mux {
	r := newRouteur()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		root(w)
	})

	r.Group(func(protected chi.Router) {
		protected.Use(middleware2.AuthMiddleware)
		protected.Route("/api/gestion", func(r chi.Router) {
			r.Post("/deposit", func(w http.ResponseWriter, r *http.Request) {
				deposit(db, w, r)
			})
			r.Post("/withdraw", func(w http.ResponseWriter, r *http.Request) {
				withdraw(db, w, r)
			})
		})
	})
	r.Get("/api/gestion/accountValue/{id}", func(w http.ResponseWriter, r *http.Request) {
		accountValue(db, w, r)
	})

	//Account Creation
	r.Route("/api/create", func(r chi.Router) {
		r.Get("/createCustomer", func(w http.ResponseWriter, r *http.Request) {
			createCustomer(db, w, r)
		})
		r.Get("/createUser", func(w http.ResponseWriter, r *http.Request) {
			createUser(db, w, r)
		})
	})

	//Account Creation
	r.Post("/api/auth/userLogin", func(w http.ResponseWriter, r *http.Request) {
		login(db, w, r)
	})

	return r
}

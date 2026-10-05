package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func accountValue(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("Utilisateur demandé : " + id))
}

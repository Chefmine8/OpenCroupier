package api

import (
	"encoding/json"
	"net/http"

	"github.com/Chefmine8/OpenCroupier/internal"
)

func root(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(internal.Response{
		Status:  "ok",
		Message: "Bienvenue sur l'API d'OpenCroupier",
	})
}

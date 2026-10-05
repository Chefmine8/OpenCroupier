package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Chefmine8/OpenCroupier/internal"
	sqlite "github.com/Chefmine8/OpenCroupier/internal/sql"
	"github.com/go-chi/chi/v5"
)

func createCustomer(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var req internal.CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	sqlite.AddCustomer(db, req.Name, req.UID)
}

func accountValue(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("Utilisateur demandé : " + id))
}

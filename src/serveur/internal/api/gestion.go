package api

import (
	"database/sql"
	"fmt"
	"net/http"

	sqlite "github.com/Chefmine8/OpenCroupier/internal/sql"
	"github.com/go-chi/chi/v5"
)

func accountValue(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	value := sqlite.CustomerValue(uid, db)
	if value == -1 {
		http.Error(w, "invalid user uid", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%v", value)
	}
}

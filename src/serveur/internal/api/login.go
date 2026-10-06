package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Chefmine8/OpenCroupier/internal"
	"github.com/Chefmine8/OpenCroupier/internal/auth"
	sqlite "github.com/Chefmine8/OpenCroupier/internal/sql"
)

func login(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var req internal.LoginRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	isUser, tokenLogin := sqlite.GetUserByUsername(db, req.UserName, req.Password)
	if !isUser {
		w.Write([]byte("Incorrect username or password"))
	} else {
		tokenStr, err := auth.NewJWT(tokenLogin.Uid, req.UserName)
		if err != nil {
			http.Error(w, "Token Generation error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(internal.LoginResponse{Token: tokenStr})
	}
}

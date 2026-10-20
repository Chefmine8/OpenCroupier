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
	req, err := ReadJSON[internal.LoginRequest](w, r)
	if err != nil {
		return
	}

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

package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Chefmine8/OpenCroupier/internal"
	sqlite "github.com/Chefmine8/OpenCroupier/internal/sql"
	"golang.org/x/crypto/bcrypt"
)

func createCustomer(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var req internal.CreateCustomerRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	lastID := sqlite.AddCustomer(db, req.Name, req.UID)
	if lastID == 0 {
		w.Write([]byte("Customer already exist"))
	} else {
		w.Write([]byte(fmt.Sprintf("Customer insert with %d ID\n", lastID)))
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func createUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	req, err := ReadJSON[internal.CreateUserRequest](w, r)
	if err != nil {
		return
	}

	hashedPass, err := hashPassword(req.Pass)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	lastID := sqlite.AddUser(db, req.UID, req.UserName, hashedPass)
	if lastID == 0 {
		w.Write([]byte("User already exist"))
	} else {
		w.Write([]byte(fmt.Sprintf("User insert with %d ID\n", lastID)))
	}
}

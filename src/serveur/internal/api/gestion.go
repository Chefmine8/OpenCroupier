package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Chefmine8/OpenCroupier/internal"
	sqlite "github.com/Chefmine8/OpenCroupier/internal/sql"
	"github.com/go-chi/chi/v5"
)

/* Aux func */
// Money aux
func isValideAmount(req internal.MoneyChange, db *sql.DB, w http.ResponseWriter) int {
	money := sqlite.MoneyChange(req.Amount, req.UID, db)
	if money == -1 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return -1
	} else if money == -2 {
		http.Error(w, "Not enough money", http.StatusBadRequest)
		return -1
	}

	return money
}

func returnMoney(req internal.MoneyChange, db *sql.DB, w http.ResponseWriter, r *http.Request) {
	money := isValideAmount(req, db, w)
	if money == -1 {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(internal.MoneyChange{
		UID:    req.UID,
		Amount: money,
	})
}

//

/* Main Func */
func accountValue(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	value := sqlite.CustomerValue(uid, db)
	if value == -1 {
		http.Error(w, "Invalid user uid", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%v", value)
	}
}

func deposit(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	req, err := ReadJSON[internal.MoneyChange](w, r)
	if err != nil {
		return
	}
	if req.Amount < 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	returnMoney(req, db, w, r)
}

func withdraw(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	req, err := ReadJSON[internal.MoneyChange](w, r)
	if err != nil {
		return
	}

	req.Amount *= -1
	returnMoney(req, db, w, r)
}

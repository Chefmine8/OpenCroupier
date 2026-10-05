package sqlite

import (
	"database/sql"
	"log"
)

func CustomerValue(uid string, db *sql.DB) int {
	var accountValue int
	err := db.QueryRow("SELECT accountValue FROM customer WHERE uid = ?", uid).Scan(&accountValue)
	if err == sql.ErrNoRows {
		return -1
	} else if err != nil {
		log.Fatalf("QueryRow Error : %v", err)
	}
	return accountValue
}

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
		return -1
	}
	return accountValue
}

func MoneyChange(value int, uid string, db *sql.DB) int {
	currentMoney := CustomerValue(uid, db)
	if currentMoney == -1 {
		return -1
	}
	money := currentMoney + value
	if money < 0 {
		return -2
	}

	_, err := db.Exec("UPDATE customer SET accountValue = ? WHERE uid = ?", money, uid)
	if err != nil {
		log.Fatalf("QueryRow Error : %v", err)
		return -1
	}

	return money
}

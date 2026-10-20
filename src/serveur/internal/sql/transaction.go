package sqlite

import (
	"database/sql"
	"log"
)

func AddTransaction(db *sql.DB, customer_uid string, user_uid string, amount int) int64 {
	balance_after := CustomerValue(customer_uid, db)
	insertCustomer := `INSERT INTO customer (customer_uid, user_uid, amount, balance_after) VALUES (?, ?, ?, ?, )`
	res, err := db.Exec(insertCustomer, customer_uid, user_uid, amount, balance_after)
	if err != nil {
		log.Printf("Insertion transaction error : %v", err)
		return 0
	}

	lastID, _ := res.LastInsertId()
	affectedRows, _ := res.RowsAffected()
	if affectedRows > 0 {
		return lastID
	} else {
		return 0
	}
	return 0
}

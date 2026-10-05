package sqlite

import (
	"database/sql"
	"fmt"
	"log"
)

func AddCustomer(db *sql.DB, name string, uid string) int64 {
	insertCustomer := `INSERT OR IGNORE INTO customer (uid, accountValue, customerName) VALUES (?, ?, ?)`
	res, err := db.Exec(insertCustomer, uid, 1050, name)
	if err != nil {
		log.Fatalf("Insertion customer error : %v", err)
	}

	lastID, _ := res.LastInsertId()
	affectedRows, _ := res.RowsAffected()
	if affectedRows > 0 {
		return lastID
	} else {
		fmt.Println()
		return 0
	}
	return 0
}

func AddUser(db *sql.DB, uid string, name string, password string) int64 {
	insertCustomer := `INSERT OR IGNORE INTO user (uid, userName, password) VALUES (?, ?, ?)`
	res, err := db.Exec(insertCustomer, uid, name, password)
	if err != nil {
		log.Fatalf("Insertion customer error : %v", err)
	}

	lastID, _ := res.LastInsertId()
	affectedRows, _ := res.RowsAffected()
	if affectedRows > 0 {
		return lastID
	} else {
		fmt.Println()
		return 0
	}
	return 0
}

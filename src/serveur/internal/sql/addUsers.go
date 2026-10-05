package sqlite

import (
	"database/sql"
	"fmt"
	"log"
)

func AddCustomer(db *sql.DB, name string, uid string) {
	insertCustomer := `INSERT OR IGNORE INTO customer (uid, accountValue, customerName) VALUES (?, ?, ?)`
	res, err := db.Exec(insertCustomer, uid, 1050, name)
	if err != nil {
		log.Fatalf("Erreur insertion : %v", err)
	}

	lastID, _ := res.LastInsertId()
	affectedRows, _ := res.RowsAffected()
	if affectedRows > 0 {
		fmt.Printf("Utilisateur inséré avec l'ID : %d\n", lastID)
	} else {
		fmt.Println("Utilisateur déjà existant (ignoré).")
	}
}

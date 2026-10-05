package sqlite

import (
	"database/sql"
	"fmt"
	"log"
)

func OpenDB() *sql.DB {
	db, err := sql.Open("sqlite", "customer.db")
	if err != nil {
		log.Fatalf("Erreur d'ouverture : %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Connexion impossible : %v", err)
	}
	fmt.Println("Base de données connectée/créée avec succès !")

	return db
}

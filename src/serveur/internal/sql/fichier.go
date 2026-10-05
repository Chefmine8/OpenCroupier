package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func createTable(db *sql.DB) {
	creerTableSQL := `
	CREATE TABLE IF NOT EXISTS customer (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uid TEXT UNIQUE NOT NULL,
		accountValue INTEGER NOT NULL,
		customerName TEXT NOT NULL
	);`

	if _, err := db.Exec(creerTableSQL); err != nil {
		log.Fatalf("Erreur création table : %v", err)
	}

	creerTableSQL = `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nom TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`

	if _, err := db.Exec(creerTableSQL); err != nil {
		log.Fatalf("Erreur création table : %v", err)
	}
}

func OpenDB() *sql.DB {
	db, err := sql.Open("sqlite", "user.db")
	if err != nil {
		log.Fatalf("Erreur d'ouverture : %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Connexion impossible : %v", err)
	}
	fmt.Println("Base de données connectée/créée avec succès !")

	createTable(db)
	return db
}

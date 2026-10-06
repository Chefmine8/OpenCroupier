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
		customerName TEXT UNIQUE NOT NULL
	);`

	if _, err := db.Exec(creerTableSQL); err != nil {
		log.Fatalf("Creation customer table error : %v", err)
	}

	creerTableSQL = `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uid TEXT UNIQUE NOT NULL,
		userName TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`

	if _, err := db.Exec(creerTableSQL); err != nil {
		log.Fatalf("Creation user table error : %v", err)
	}
}

func OpenDB() *sql.DB {
	db, err := sql.Open("sqlite", "user.db")
	if err != nil {
		log.Fatalf("Opened error : %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Impossible connexion to the db : %v", err)
	}
	fmt.Println("DB successfully opened !")

	createTable(db)
	return db
}

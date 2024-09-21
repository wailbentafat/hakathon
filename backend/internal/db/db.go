package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

var db *sql.DB

func InitDb(dsn string) (*sql.DB, error) {
	log.Printf("Initializing database connection with dsn: %s", dsn)
	var err error
	db, err = sql.Open("sqlite3", dsn) // Use "sqlite3" for SQLite
	if err != nil {
		log.Printf("Error opening database connection: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, err
	}

	// Create tables
	createtablestuff := `CREATE TABLE IF NOT EXISTS staff (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		password TEXT NOT NULL,
		email TEXT NOT NULL,
		is_admin BOOLEAN NOT NULL DEFAULT false,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL
	);`
	
	createtablecomplaints := `CREATE TABLE IF NOT EXISTS complaint (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bank_card TEXT,
		name TEXT NOT NULL,
		category TEXT NOT NULL,
		location TEXT,
		phone_number TEXT,
		description TEXT,
		staff_id INTEGER,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		email_address TEXT,
		bank_name TEXT,
		website_url TEXT,
		national_id_number TEXT,
		card_type TEXT,
		incident_date DATE,
		transaction_amount DECIMAL(10, 2),
		transaction_date DATE,
		merchant_name TEXT,
		merchant_registration TEXT,
		FOREIGN KEY (staff_id) REFERENCES staff(id) ON DELETE SET NULL
	);`

	createcategorie := `CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);`

	// Execute table creation queries
	_, err = db.Exec(createtablestuff)
	if err != nil {
		log.Printf("Error creating staff table: %v", err)
		return nil, err
	}

	_, err = db.Exec(createtablecomplaints)
	if err != nil {
		log.Printf("Error creating complaints table: %v", err)
		return nil, err
	}

	_, err = db.Exec(createcategorie)
	if err != nil {
		log.Printf("Error creating categories table: %v", err)
		return nil, err
	}

	log.Println("Database connection initialized successfully")
	return db, nil
}

func GetDb() *sql.DB {
	return db
}

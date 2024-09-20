package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDb(dsn string) (*sql.DB, error) {
    log.Printf("Initializing database connection with dsn: %s", dsn)
    var err error
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Printf("Error opening database connection: %v", err)
        return nil, err
    }

    if err := db.Ping(); err != nil {
        log.Printf("Error pinging database: %v", err)
        return nil, err
    }
    fmt.Println(db)
	createtablestuff := `CREATE TABLE IF NOT EXISTS staff (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		is_admin BOOLEAN NOT NULL DEFAULT false,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL
	);`
	createtablecomplaints := `CREATE TABLE IF NOT EXISTS complaints (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		bank_card VARCHAR(255),
		name VARCHAR(255) NOT NULL,
		category VARCHAR(255) NOT NULL,
		location VARCHAR(255),
		phone_number VARCHAR(255),
		description VARCHAR(255),
		staff_id BIGINT UNSIGNED,
		FOREIGN KEY (staff_id) REFERENCES staff(id) ON DELETE SET NULL,
		satisfied BOOLEAN,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	
	createcategorie := `CREATE TABLE IF NOT EXISTS categories (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL
	);`
	_, err = db.Exec(createtablestuff)
	if err != nil {
		log.Printf("Error creating users table: %v", err)
		return nil, err
	}
   

    _, err = db.Exec(createtablecomplaints)
    if err != nil {
        log.Printf("Error creating users table: %v", err)
        return nil, err
    }

    _, err = db.Exec(createcategorie)
    if err != nil {
        log.Printf("Error creating users_profiles table: %v", err)
        return nil, err
    }
	
    
    log.Println("Database connection initialized successfully")
    return db, nil
}

func GetDb() *sql.DB {
    return db
}

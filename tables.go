package main

import (
	"database/sql"
	"log"
)

func initializeTables(db *sql.DB) {
	InfoLogger.Println("Setting up base tables...")
	users := `DROP TABLE IF EXISTS users;
	CREATE TABLE users (
		id 			SERIAL PRIMARY KEY,
		name		TEXT NOT NULL,
		surname		TEXT NOT NULL,
		age			INT
	);
	`
	executeSQL(db, users)
	books := `DROP TABLE IF EXISTS books;
	CREATE TABLE books (
		id 			SERIAL PRIMARY KEY,
		ISBN		INT UNIQUE,
		title		TEXT NOT NULL,
		author		TEXT NOT NULL,
		pages		INT
	);
	`
	executeSQL(db, books)
	authors := `DROP TABLE IF EXISTS authors;
	CREATE TABLE authors (
		id 			SERIAL PRIMARY KEY,
		ISBN		INT UNIQUE NOT NULL,
		title		TEXT NOT NULL,
		author		TEXT NOT NULL,
		pages		INT
	);
	`
	executeSQL(db, authors)
}

func executeSQL(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		ErrorLogger.Printf("Failed to execute query, following error occured: %s\n", err.Error())
		log.Fatal(err)
	}
}

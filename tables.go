package main

import (
	"database/sql"
	"fmt"
	"log"
)

func initializeTables() {
	InfoLogger.Println("Setting up base tables...")
	users := `DROP TABLE IF EXISTS users;
	CREATE TABLE users (
		id 			SERIAL PRIMARY KEY,
		name		TEXT NOT NULL,
		surname		TEXT NOT NULL,
		age			INT
	);
	`
	executeSQL(DB, users)
	executeSQL(DB, fmt.Sprintf("INSERT INTO users (name, surname, age) VALUES ('wiaan', 'botha', 23);"))
	books := `DROP TABLE IF EXISTS books;
	CREATE TABLE books (
		id 			SERIAL PRIMARY KEY,
		ISBN		INT UNIQUE,
		title		TEXT NOT NULL,
		author		TEXT NOT NULL,
		pages		INT
	);
	`
	executeSQL(DB, books)
	executeSQL(DB, fmt.Sprintf("INSERT INTO books (ISBN, title, author, pages) VALUES (2341, 'The Martian Threat', 'Dave Campbell', 342);"))
	authors := `DROP TABLE IF EXISTS authors;
	CREATE TABLE authors (
		id 			SERIAL PRIMARY KEY,
		name		TEXT NOT NULL,
		books		INT,
		awards		INT
	);
	`
	executeSQL(DB, authors)
	executeSQL(DB, fmt.Sprintf("INSERT INTO authors (name, books, awards) VALUES ('Dave Campbell', 1, 0);"))
}

func executeSQL(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		ErrorLogger.Printf("Failed to execute query, following error occured: %s\n", err.Error())
		log.Fatal(err)
	}
}

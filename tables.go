package main

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
	executeSQL(users)
	executeSQL("INSERT INTO users (name, surname, age) VALUES ('wiaan', 'botha', 23);")
	books := `DROP TABLE IF EXISTS books;
	CREATE TABLE books (
		id 			SERIAL PRIMARY KEY,
		ISBN		INT UNIQUE,
		title		TEXT NOT NULL,
		author		TEXT NOT NULL,
		pages		INT
	);
	`
	executeSQL(books)
	executeSQL("INSERT INTO books (ISBN, title, author, pages) VALUES (2341, 'The Martian Threat', 'Dave Campbell', 342);")
	authors := `DROP TABLE IF EXISTS authors;
	CREATE TABLE authors (
		id 			SERIAL PRIMARY KEY,
		name		TEXT NOT NULL,
		books		INT,
		awards		INT
	);
	`
	executeSQL(authors)
	executeSQL("INSERT INTO authors (name, books, awards) VALUES ('Dave Campbell', 1, 0);")
}

func executeSQL(query string) {
	_, err := DB.Exec(query)
	errorHandle(err, "Failed to execute query")
}

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
}

func executeSQL(query string) {
	_, err := DB.Exec(query)
	errorHandle(err, "Failed to execute query")
}

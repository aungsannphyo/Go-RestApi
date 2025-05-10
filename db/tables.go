package db

func createUserTable() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email VARCHAR(255) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL
	)
	`
	_, err := DBInstance.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table")
	}
}

func createEventsTable() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(255) NOT NULL,
	description TEXT NOT NULL,
	location VARCHAR(255) NOT NULL,
	dateTime DATETIME NOT NULL,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err := DBInstance.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table")
	}
}

func createTable() {
	createUserTable()
	createEventsTable()
}

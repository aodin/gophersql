package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var createUsers = `
CREATE TABLE users (
    id INTEGER,
    name VARCHAR
);
`

var insertUser = `INSERT INTO users (id, name) VALUES (?, ?)`

var selectUser = `SELECT id, name FROM users`

type User struct {
	ID   int64
	Name sql.NullString
}

func (u User) String() string {
	if u.Name.Valid {
		return u.Name.String
	}
	return "No name"
}

func main() {
	// Connect to an in-memory sqlite3 instance
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create the table
	if _, err := db.Exec(createUsers); err != nil {
		panic(err)
	}

	// Insert a user without a name
	if _, err := db.Exec(insertUser, 1, nil); err != nil {
		panic(err)
	}

	// Select a user
	var user User
	row := db.QueryRow(selectUser)
	if err = row.Scan(&user.ID, &user.Name); err != nil {
		panic(err)
	}
	log.Println(user)
}

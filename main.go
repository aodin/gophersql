package main

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var createUsers = `
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name VARCHAR
);
`

var insertUser = `INSERT INTO users (id, name) VALUES (?, ?)`

var selectUser = `SELECT id, name FROM users`

func main() {
    
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    if _, err := db.Exec(createUsers); err != nil {
        panic(err)
    }

    if _, err := db.Exec(insertUser, 23, "skidoo"); err != nil {
        panic(err)
    }

    var id int64
    var name string
    row := db.QueryRow(selectUser)

    if err = row.Scan(&id, &name); err != nil {
        panic(err)
    }

    log.Println(id, name)





}
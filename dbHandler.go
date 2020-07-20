package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func addToDB(token string) {
	database, _ := sql.Open("sqlite3", "./tokens.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS tokens (id INTEGER PRIMARY KEY, token TEXT, timestamp TEXT)")
	statement.Exec()

	statement, _ = database.Prepare("INSERT INTO people (token, timestamp) VALUES (?, ?)")
	statement.Exec(token, time.Now())
}

func checkIfExists(token string) { //TODO prevent SQL injection
	database, _ := sql.Open("sqlite3", "./tokens.db")
	rows, _ := database.Query("SELECT id, token, timestamp FROM tokens")
	var id int
	var tokenFromDB string
	var timestampFromDB string
	for rows.Next() {
		rows.Scan(&id, &tokenFromDB, &timestampFromDB)
		fmt.Println(strconv.Itoa(id) + ": " + tokenFromDB + " " + timestampFromDB)
	}
}

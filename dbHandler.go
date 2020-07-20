package main

import (
	"database/sql"
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

func checkIfExists(token string) {
	/*rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)*/
}

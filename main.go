package main

import (
	"calendar/template"
	"calendar/util"

	"database/sql"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// init DB
	db, err := sql.Open("sqlite3", "db/main.db")

	util.Printerr(err)

	defer db.Close()

	sts := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY,
		name TEXT, 
		passwd TEXT
	);`

	_, err = db.Exec(sts)

	util.Printerr(err)

	fmt.Println("DB was initialized.")

	// init templ
	component := template.Index()

	// init http server
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on port 3000.")
	http.ListenAndServe(":3000", nil)
}

package helpers

import (
	"database/sql"
	"fmt"
)

// QueryDb - Execute sql statements through helper
func QueryDb(db *sql.DB, query string, args []string) {
	// Create table query
	stmt, err := db.Prepare(query)
	CheckErr(err, "")

	_, err = stmt.Exec(args)
	// Exit if something goes wrong with our SQL statement above
	CheckErr(err, "")
}

// InitDB - Initialise DB connection
func InitDB(filepath string) {
	db, err := sql.Open("sqlite3", filepath)
	CheckErr(err, "")

	// Create table query
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS users( id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,fname VARCHAR NOT NULL, lname VARCHAR NOT NULL, password VARCHAR NOT NULL, username VARCHAR NOT NULL);")
	CheckErr(err, "")

	_, err = stmt.Exec()
	// Exit if something goes wrong with our SQL statement above
	CheckErr(err, "")

	fmt.Println("Created Users Table.")

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
}

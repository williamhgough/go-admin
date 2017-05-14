package helpers

import (
	"database/sql"
	"fmt"
)

// QueryDb - Execute sql statements through helper
func QueryDb(db *sql.DB, method string, query string, args map[string]string) sql.Result {
	// Create table query
	stmt, err := db.Prepare(query)
	CheckErr(err, "")

	// Create values interface slice
	values := make([]interface{}, 0, len(args))

	// Check for args
	if args != nil {
		for key := range args {
			values = append(values, args[key])
		}
		res, err := stmt.Exec(values...)
		// Exit if something goes wrong with our SQL statement above
		CheckErr(err, "")

		// Log successful query
		fmt.Printf("%s query completed.\n", method)

		return res
	}

	res, err := stmt.Exec()
	// Exit if something goes wrong with our SQL statement above
	CheckErr(err, "")

	// Log successful query
	fmt.Printf("%s query completed.\n", method)

	return res

}

// InitDB - Initialise DB connection
func InitDB(filepath string) {
	db, err := sql.Open("sqlite3", filepath)
	CheckErr(err, "Can't open the database.")

	// Create Users Table
	QueryDb(db, "Create", "CREATE TABLE IF NOT EXISTS users( id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,fname VARCHAR NOT NULL, lname VARCHAR NOT NULL, username VARCHAR NOT NULL, password VARCHAR NOT NULL);", nil)
	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
}

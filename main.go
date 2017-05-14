package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/williamhgough/go-admin/helpers"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	// Create new database instance
	helpers.InitDB("storage.db")

	// Create new instance of echo
	iris := iris.New()

	// Add HTTPRouter
	iris.Adapt(httprouter.New())

	// Gzip all responses
	iris.Config.Gzip = true

	// Serve up build dir
	iris.StaticWeb("/", "build/")

	// start server on port:8080
	iris.Listen(":8080")
}

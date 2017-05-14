package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/williamhgough/go-admin/helpers"
)

func main() {
	// Create new database instance
	helpers.InitDB("storage.db")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/", "build")

	e.Start(":8080")
}

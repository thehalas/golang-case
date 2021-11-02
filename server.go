package main

import (
	"github.com/labstack/echo/v4"
	"vatansoft-golang-case/database"
	"vatansoft-golang-case/routes"
)

func main() {
	database.Connect()
	e := echo.New()
	routes.Setup(e)
	e.Logger.Fatal(e.Start("localhost:8080"))
}

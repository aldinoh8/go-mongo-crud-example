package main

import (
	"example/config"
	"example/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.InitDB()
	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}

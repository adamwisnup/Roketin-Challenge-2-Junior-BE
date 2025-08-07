package main

import (
	"challenge-2/config"
	"challenge-2/models"
	"challenge-2/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.ConnectDB()
	config.DB.AutoMigrate(&models.Movie{})

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}

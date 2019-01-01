package main

import (
	"os"
	"unishare/db"
	"unishare/routes"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.InitRoutes(e)
	db.CreateTables()

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}

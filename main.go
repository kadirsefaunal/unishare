package main

import (
	"os"
	"unishare/db"
	"unishare/routes"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	port := os.Getenv("PORT")

	routes.InitRoutes(e)
	db.CreateTables()

	e.Logger.Fatal(e.Start(":" + port))
}

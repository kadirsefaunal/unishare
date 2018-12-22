package main

import (
	"os"
	"unishare/routes"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	port := os.Getenv("PORT")

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":" + port))
}

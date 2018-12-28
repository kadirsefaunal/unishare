package db

import (
	"os"
	"unishare/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Connect() *gorm.DB {
	dbType := os.Getenv("DB_TYPE")
	dbName := os.Getenv("DB_NAME")

	db, err := gorm.Open(dbType, dbName)
	if err != nil {
		panic(err)
	}

	return db
}

func CreateTables() {
	db := Connect()
	defer db.Close()

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.School{})
}

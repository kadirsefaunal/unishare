package repositories

import (
	"unishare/db"
	"unishare/models"
)

func SchoolInsert(school *models.School) error {
	db := db.Connect()
	defer db.Close()

	err := db.Create(&school).Error
	return err
}

func SchoolDelete(school models.School) error {
	db := db.Connect()
	defer db.Close()

	err := db.Delete(school).Error
	return err
}

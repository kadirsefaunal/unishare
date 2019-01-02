package repositories

import (
	"unishare/db"
	"unishare/models"
)

func ClassInsert(class *models.Class) error {
	db := db.Connect()
	defer db.Close()

	err := db.Create(&class).Error
	return err
}

func ClassDelete(class models.Class) error {
	db := db.Connect()
	defer db.Close()

	err := db.Delete(class).Error
	return err
}

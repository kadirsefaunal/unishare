package repositories

import (
	"unishare/db"
	"unishare/models"
)

func PostInsert(post *models.Post) error {
	db := db.Connect()
	defer db.Close()

	err := db.Create(&post).Error
	return err
}

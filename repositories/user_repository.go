package repositories

import (
	"unishare/db"
	"unishare/models"
)

func UserInsert(user *models.User) error {
	db := db.Connect()
	defer db.Close()

	err := db.Create(&user).Error
	return err
}

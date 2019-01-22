package repositories

import (
	"unishare/db"
	"unishare/models"
)

func UserGetByUserName(userName string) (*models.User, error) {
	db := db.Connect()
	defer db.Close()

	user := new(models.User)
	err := db.Where("user_name = ?", userName).Find(&user).Error

	return user, err
}

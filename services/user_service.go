package services

import (
	"unishare/helpers"
	"unishare/models"
	"unishare/repositories"
)

func UserCreate(user *models.User) error {
	user.Password = helpers.EncryptPassword(user.Password)

	return repositories.UserInsert(user)
}

package services

import (
	"errors"
	"unishare/helpers"
	"unishare/models"
	"unishare/repositories"
)

func UserCreate(user *models.User) error {
	user.Password = helpers.EncryptPassword(user.Password)

	return repositories.UserInsert(user)
}

func UserLogin(loginUser *models.LoginUser) error {
	user, err := repositories.UserGetByUserName(loginUser.UserName)
	if err != nil {
		return err
	}

	err = helpers.CheckPassword(user.Password, loginUser.Password)
	if err != nil {
		return errors.New("password is not correct")
	}

	return nil
}

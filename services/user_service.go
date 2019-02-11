package services

import (
	"errors"
	"unishare/helpers"
	"unishare/models"
	"unishare/repositories"
	"unishare/structs"
)

func UserCreate(user *models.User) (*structs.Token, error) {
	user.Password = helpers.EncryptPassword(user.Password)

	err := repositories.Insert(user)
	if err != nil {
		return nil, err
	}

	return RedisSetUser(user), nil
}

func UserLogin(loginUser *models.LoginUser) (*structs.Token, error) {
	user, err := repositories.UserGetByUserName(loginUser.UserName)
	if err != nil {
		return nil, err
	}

	err = helpers.CheckPassword(user.Password, loginUser.Password)
	if err != nil {
		return nil, errors.New("Password is not correct")
	}

	token := RedisSetUser(user)

	return token, nil
}

func GetCurrentUser(token string) (*models.User, error) {
	return RedisGetUser(token)
}

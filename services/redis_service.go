package services

import (
	"encoding/json"
	"fmt"
	"time"
	"unishare/helpers"
	"unishare/models"
	"unishare/storages"
	"unishare/structs"
)

func RedisSetUser(user *models.User) *structs.Token {
	token := new(structs.Token)
	token.UserID = user.ID
	token.UserName = user.UserName
	token.AccessToken = helpers.CreateTokenByUserName(user.UserName)

	jsonUser, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	client := storages.GetRedisClient()
	client.Set(token.AccessToken, jsonUser, time.Hour)

	return token
}

func RedisCheckUser(token string) bool {
	client := storages.GetRedisClient()

	user, err := client.Get(token).Result()
	if err != nil {
		return false
	}

	return (user != "")
}

func RedisGetUser(token string) (*models.User, error) {
	client := storages.GetRedisClient()

	cacheUser, err := client.Get(token).Result()
	if err != nil {
		fmt.Println(err.Error())
	}

	user := new(models.User)

	if err := json.Unmarshal([]byte(cacheUser), &user); err != nil {
		return user, err
	}

	return user, nil
}

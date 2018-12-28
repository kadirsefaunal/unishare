package services

import (
	"encoding/json"
	"time"
	"unishare/helpers"
	"unishare/models"
	"unishare/storages"
)

func RedisSetUser(user *models.User) string {
	token := helpers.CreateTokenByUserName(user.UserName)

	jsonUser, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	client := storages.GetRedisClient()
	client.Set(token, jsonUser, time.Hour)

	return token
}

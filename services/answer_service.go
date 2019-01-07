package services

import (
	"unishare/models"
	"unishare/repositories"
)

func AnswerInsert(answer *models.Answer, token string) error {
	user, err := GetCurrentUser(token)
	if err != nil {
		return err
	}

	answer.User = user
	return repositories.AnswerInsert(answer)
}

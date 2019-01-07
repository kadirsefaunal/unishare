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

func AnswerGet(answerID string) (*models.Answer, error) {
	return repositories.AnswerGet(answerID)
}

func AnswerList(token string) (*[]models.Answer, error) {
	user, err := GetCurrentUser(token)
	if err != nil {
		return nil, err
	}

	return repositories.AnswerList(user.ID)
}

func AnswerUpdate(answer *models.Answer) error {
	return repositories.AnswerUpdate(answer)
}

func AnswerDelete(answerID string) error {
	answer, err := AnswerGet(answerID)
	if err != nil {
		return err
	}

	return repositories.AnswerDelete(answer)
}

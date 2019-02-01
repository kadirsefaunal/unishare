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
	return repositories.Insert(answer)
}

func AnswerGet(answerID string) (interface{}, error) {
	return repositories.FindByID(answerID, new(models.Answer))
}

func AnswerList(token string) (interface{}, error) {
	user, err := GetCurrentUser(token)
	if err != nil {
		return nil, err
	}

	return repositories.GetListByUserID(user.ID, new([]models.Answer))
}

func AnswerUpdate(answer interface{}) error {
	return repositories.Update(answer)
}

func AnswerDelete(answerID string) error {
	answer, err := AnswerGet(answerID)
	if err != nil {
		return err
	}

	return repositories.Delete(answer)
}

package services

import (
	"unishare/helpers"
	"unishare/models"
	"unishare/repositories"
	"unishare/structs"
)

func AnswerInsert(answer *models.Answer, token string, postID string) (*structs.Info, error) {
	user, err := GetCurrentUser(token)
	if err != nil {
		return nil, err
	}

	answer.User = user
	answer.QuestionID = helpers.StringToUint(postID)

	err = repositories.Insert(answer)
	if err != nil {
		return nil, err
	}

	info := new(structs.Info)
	info.ID = answer.ID
	info.Status = "created"

	return info, nil
}

func AnswerGet(answerID string) (interface{}, error) {
	return repositories.FindByID(answerID, new(models.Answer))
}

func AnswerList(questionID string) (interface{}, error) {
	return repositories.GetAnswerByQuestionID(questionID)
}

func AnswerUpdate(answer interface{}) (*structs.Info, error) {
	err := repositories.Update(answer)
	if err != nil {
		return nil, err
	}

	info := new(structs.Info)
	info.ID = (answer.(*models.Answer)).ID
	info.Status = "updated"

	return info, nil
}

func AnswerDelete(answerID string) (*structs.Info, error) {
	answer, err := AnswerGet(answerID)
	if err != nil {
		return nil, err
	}

	err = repositories.Delete(answer)
	if err != nil {
		return nil, err
	}

	info := new(structs.Info)
	info.ID = helpers.StringToUint(answerID)
	info.Status = "deleted"

	return info, nil
}

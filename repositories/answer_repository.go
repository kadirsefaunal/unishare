package repositories

import (
	"unishare/db"
	"unishare/models"
)

func AnswerInsert(answer *models.Answer) error {
	db := db.Connect()
	defer db.Close()

	err := db.Create(&answer).Error
	return err
}

func AnswerGet(answerID string) (*models.Answer, error) {
	db := db.Connect()
	defer db.Close()

	answer := new(models.Answer)
	err := db.Find(&answer, "id = ?", answerID).Error

	return answer, err
}

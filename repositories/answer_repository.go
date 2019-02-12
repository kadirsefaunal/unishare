package repositories

import (
	"unishare/db"
	"unishare/models"
)

func GetAnswerByQuestionID(questionID string) (*[]models.Answer, error) {
	db := db.Connect()
	defer db.Close()

	answers := new([]models.Answer)

	err := db.Find(answers, "question_id = ?", questionID).Error

	return answers, err
}

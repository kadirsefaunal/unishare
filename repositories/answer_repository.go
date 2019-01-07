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

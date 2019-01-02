package services

import (
	"unishare/models"
	"unishare/repositories"
)

func ClassCreate(class *models.Class) error {
	return repositories.ClassInsert(class)
}

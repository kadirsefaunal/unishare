package services

import (
	"strconv"
	"unishare/models"
	"unishare/repositories"

	"github.com/jinzhu/gorm"
)

func ClassCreate(class *models.Class) error {
	return repositories.ClassInsert(class)
}

func ClassDelete(id string) error {
	classID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	class := models.Class{Model: gorm.Model{ID: uint(classID)}}
	return repositories.ClassDelete(class)
}

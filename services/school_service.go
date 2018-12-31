package services

import (
	"strconv"
	"unishare/models"
	"unishare/repositories"

	"github.com/jinzhu/gorm"
)

func SchoolCreate(school *models.School) error {
	return repositories.SchoolInsert(school)
}

func SchoolDelete(id string) error {
	schoolID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	school := models.School{Model: gorm.Model{ID: uint(schoolID)}}

	return repositories.SchoolDelete(school)
}

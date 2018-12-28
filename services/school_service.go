package services

import (
	"unishare/models"
	"unishare/repositories"
)

func SchoolCreate(school *models.School) error {
	return repositories.SchoolInsert(school)
}

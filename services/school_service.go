package services

import (
	"unishare/models"
	"unishare/repositories"
)

func SchoolCreate(school *models.School) error {
	return repositories.Insert(school)
}

func SchoolDelete(id string) error {
	school, err := SchoolGet(id)
	if err != nil {
		return err
	}

	return repositories.Delete(school)
}

func SchoolList() (interface{}, error) {
	return repositories.GetList(new([]models.School))
}

func SchoolGet(id string) (interface{}, error) {
	return repositories.FindByID(id, new(models.School))
}

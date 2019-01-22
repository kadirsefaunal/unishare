package services

import (
	"unishare/models"
	"unishare/repositories"
)

func ClassCreate(class *models.Class) error {
	return repositories.Insert(class)
}

func ClassDelete(id string) error {
	class, err := ClassGet(id)
	if err != nil {
		return err
	}

	return repositories.Delete(class)
}

func ClassGet(id string) (interface{}, error) {
	return repositories.FindByID(id, new(models.Class))
}

func ClassUpdate(class interface{}) error {
	return repositories.Update(class)
}

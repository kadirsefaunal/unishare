package services

import (
	"unishare/models"
	"unishare/repositories"
	"unishare/structs"
)

func SchoolCreate(school *models.School) (*structs.Info, error) {
	err := repositories.Insert(school)
	if err != nil {
		return nil, err
	}

	info := new(structs.Info)
	info.ID = school.ID
	info.Status = "created"

	return info, nil
}

func SchoolDelete(id string) (*structs.Info, error) {
	school, err := SchoolGet(id)
	if err != nil {
		return nil, err
	}

	err = repositories.Delete(school)
	if err != nil {
		return nil, err
	}

	info := new(structs.Info)
	info.ID = (school.(models.School)).ID
	info.Status = "deleted"

	return info, nil
}

func SchoolList() (interface{}, error) {
	return repositories.GetList(new([]models.School))
}

func SchoolGet(id string) (interface{}, error) {
	return repositories.FindByID(id, new(models.School))
}

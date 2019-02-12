package services

import (
	"unishare/helpers"
	"unishare/models"
	"unishare/repositories"
	"unishare/structs"
)

func ClassCreate(schoolId string, class *models.Class) (*structs.Info, error) {
	class.SchoolID = helpers.StringToUint(schoolId)

	err := repositories.Insert(class)
	if err != nil {
		return nil, err
	}

	info := new(structs.Info)
	info.ID = class.ID
	info.Status = "created"

	return info, nil
}

func ClassDelete(id string) (*structs.Info, error) {
	class, err := ClassGet(id)
	if err != nil {
		return nil, err
	}

	err = repositories.Delete(class)
	if err != nil {
		return nil, err
	}

	info := new(structs.Info)
	info.ID = (class.(models.Class)).ID
	info.Status = "deleted"

	return info, nil
}

func ClassGet(id string) (interface{}, error) {
	return repositories.FindByID(id, new(models.Class))
}

func ClassUpdate(class interface{}) (*structs.Info, error) {
	err := repositories.Update(class)
	if err != nil {
		return nil, err
	}

	info := new(structs.Info)
	info.ID = (class.(models.Class)).ID
	info.Status = "updated"

	return info, nil
}

func ClassList(schoolID string) (interface{}, error) {
	return repositories.GetClassBySchoolID(schoolID)
}

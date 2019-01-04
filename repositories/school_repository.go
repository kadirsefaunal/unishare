package repositories

import (
	"unishare/db"
	"unishare/models"
)

func SchoolInsert(school *models.School) error {
	db := db.Connect()
	defer db.Close()

	err := db.Create(&school).Error
	return err
}

func SchoolDelete(school models.School) error {
	db := db.Connect()
	defer db.Close()

	err := db.Delete(school).Error
	return err
}

func SchoolList() (*[]models.School, error) {
	db := db.Connect()
	defer db.Close()

	schools := new([]models.School)

	err := db.Find(&schools).Error
	if err != nil {
		return nil, err
	}

	return schools, nil
}

func SchoolGet(id string) (*models.School, error) {
	db := db.Connect()
	defer db.Close()

	school := new(models.School)
	err := db.Find(&school, "id = ?", id).Error
	db.Model(&school).Related(&school.Classes)

	return school, err
}

package repositories

import (
	"unishare/db"
	"unishare/models"
)

func GetClassBySchoolID(schoolID string) (*[]models.Class, error) {
	db := db.Connect()
	defer db.Close()

	classes := new([]models.Class)

	err := db.Find(classes, "school_id = ?", schoolID).Error

	return classes, err
}

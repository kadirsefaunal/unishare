package repositories

import (
	"unishare/db"
	"unishare/models"
)

func ClassInsert(class *models.Class) error {
	db := db.Connect()
	defer db.Close()

	err := db.Create(&class).Error
	return err
}

func ClassDelete(class models.Class) error {
	db := db.Connect()
	defer db.Close()

	err := db.Delete(class).Error
	return err
}

func ClassGet(classID string) (*models.Class, error) {
	db := db.Connect()
	defer db.Close()

	class := new(models.Class)

	err := db.Find(&class, "id = ?", classID).Error
	return class, err
}

func ClassUpdate(class *models.Class) error {
	db := db.Connect()
	defer db.Close()

	err := db.Save(&class).Error
	return err
}

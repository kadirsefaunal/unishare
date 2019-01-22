package repositories

import (
	"unishare/db"
)

func Insert(model interface{}) error {
	db := db.Connect()
	defer db.Close()

	return db.Create(model).Error
}

func Delete(model interface{}) error {
	db := db.Connect()
	defer db.Close()

	return db.Delete(model).Error
}

func Update(model interface{}) error {
	db := db.Connect()
	defer db.Close()

	return db.Save(model).Error
}

func FindByID(id string, model interface{}) (interface{}, error) {
	db := db.Connect()
	defer db.Close()

	err := db.Find(model, "id = ?", id).Error

	return model, err
}

func GetListByUserID(userID uint, model interface{}) (interface{}, error) {
	db := db.Connect()
	defer db.Close()

	err := db.Find(model, "user_id = ?", userID).Error

	return model, err
}

func GetList(model interface{}) (interface{}, error) {
	db := db.Connect()
	defer db.Close()

	err := db.Find(model).Error

	return model, err
}

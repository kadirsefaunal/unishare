package repositories

import (
	"unishare/db"
	"unishare/models"
)

func PostInsert(post *models.Post) error {
	db := db.Connect()
	defer db.Close()

	err := db.Create(&post).Error
	return err
}

func PostGet(postID string) (*models.Post, error) {
	db := db.Connect()
	defer db.Close()

	post := new(models.Post)
	err := db.Find(&post, "id = ?", postID).Error

	return post, err
}

func PostUpdate(post *models.Post) error {
	db := db.Connect()
	defer db.Close()

	err := db.Save(&post).Error
	return err
}

func PostDelete(post *models.Post) error {
	db := db.Connect()
	defer db.Close()

	err := db.Delete(&post).Error
	return err
}

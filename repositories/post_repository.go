package repositories

import (
	"unishare/db"
	"unishare/models"
)

func GetPostByClassID(classId string) (*[]models.Post, error) {
	db := db.Connect()
	defer db.Close()

	posts := new([]models.Post)

	err := db.Find(posts, "class_id = ?", classId).Error

	return posts, err
}

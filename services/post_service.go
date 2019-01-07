package services

import (
	"unishare/models"
	"unishare/repositories"
)

func PostCreate(post *models.Post, token string) error {
	user, err := GetCurrentUser(token)
	if err != nil {
		return err
	}

	post.User = user
	return repositories.PostInsert(post)
}

func PostGet(postID string) (*models.Post, error) {
	return repositories.PostGet(postID)
}

func PostUpdate(post *models.Post) error {
	return repositories.PostUpdate(post)
}

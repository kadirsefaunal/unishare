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
	return repositories.Insert(post)
}

func PostGet(postID string) (interface{}, error) {
	return repositories.FindByID(postID, new(models.Post))
}

func PostUpdate(post interface{}) error {
	return repositories.Update(post)
}

func PostDelete(postID string) error {
	post, err := PostGet(postID)
	if err != nil {
		return err
	}

	return repositories.Delete(post)
}

func PostList(token string) (interface{}, error) {
	user, err := GetCurrentUser(token)
	if err != nil {
		return nil, err
	}

	return repositories.GetListByUserID(user.ID, new([]models.Post))
}

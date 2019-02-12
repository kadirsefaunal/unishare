package services

import (
	"unishare/helpers"
	"unishare/models"
	"unishare/repositories"
	"unishare/structs"
)

func PostCreate(post *models.Post, token string, classID string) (*structs.Info, error) {
	user, err := GetCurrentUser(token)
	if err != nil {
		return nil, err
	}

	post.User = user
	post.ClassID = helpers.StringToUint(classID)

	err = repositories.Insert(post)
	if err != nil {
		return nil, err
	}

	info := new(structs.Info)
	info.ID = post.ID
	info.Status = "created"

	return info, nil
}

func PostGet(postID string) (interface{}, error) {
	return repositories.FindByID(postID, new(models.Post))
}

func PostUpdate(post interface{}) (*structs.Info, error) {
	err := repositories.Update(post)
	if err != nil {
		return nil, err
	}

	info := new(structs.Info)
	info.ID = (post.(models.Post)).ID
	info.Status = "updated"

	return info, err
}

func PostDelete(postID string) (*structs.Info, error) {
	post, err := PostGet(postID)
	if err != nil {
		return nil, err
	}

	err = repositories.Delete(post)
	if err != nil {
		return nil, err
	}

	info := new(structs.Info)
	info.ID = (post.(models.Post)).ID
	info.Status = "deleted"

	return info, nil
}

func PostList(classID string) (interface{}, error) {
	return repositories.GetPostByClassID(classID)
}

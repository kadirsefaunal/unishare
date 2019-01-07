package models

import (
	"github.com/jinzhu/gorm"
)

type Answer struct {
	gorm.Model
	Content    string `json:"content" gorm:"size:1000"`
	User       *User  `gorm:"foreignkey:UserID"`
	UserID     uint   `json:"userId"`
	Question   *Post  `gorm:"foreignkey:QuestionID"`
	QuestionID uint   `json:"questionId"`
}

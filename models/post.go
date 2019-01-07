package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title   string `json:"title" gorm:"size:200"`
	Content string `json:"content" gorm:"size:1000"`
	Type    string `json:"type" gorm:"size:50"` // Announcement/Question
	User    *User  `gorm:"foreignkey:UserID"`
	UserID  uint   `json:"userId"`
	Class   Class  `gorm:"foreignkey:ClassID"`
	ClassID uint   `json:"classId"`
}

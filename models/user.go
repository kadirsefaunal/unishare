package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	UserName      string `json:"username" gorm:"size:255"`
	Password      string `json:"password" gorm:"size:255"`
	StudentNumber string `json:"studentNumber" gorm:"size:20"`
	NameSurname   string `json:"name" gorm:"size:255"`
	Mail          string `json:"mail" gorm:"size:255"`
	StudentType   string `json:"studentType" gorm:"size:20"`
	Type          string `json:"type" gorm:"size:20"` // Student/Instructer
}

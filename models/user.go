package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	UserName      string   `json:"username" gorm:"size:255;UNIQUE"`
	Password      string   `json:"password" gorm:"size:255"`
	StudentNumber string   `json:"studentNumber" gorm:"size:20;UNIQUE"`
	NameSurname   string   `json:"name" gorm:"size:255"`
	Mail          string   `json:"mail" gorm:"size:255;UNIQUE"`
	StudentType   string   `json:"studentType" gorm:"size:20"`
	Type          string   `json:"type" gorm:"size:20"` // Student/Instructor
	School        School   `gorm:"foreignkey:SchoolID"`
	SchoolID      uint     `json:"schoolId"`
	Classes       []*Class `gorm:"many2many:user_classes;"`
}

type LoginUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

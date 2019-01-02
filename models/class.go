package models

import (
	"github.com/jinzhu/gorm"
)

type Class struct {
	gorm.Model
	Code        string `json:"code" gorm:"size:50"`
	Name        string `json:"name" gorm:"size:255"`
	Description string `json:"description" gorm:"size:1500"`
	Semester    string `json:"semester" gorm:"size:20"`
	School      School `gorm:"foreignkey:SchoolID"`
	SchoolID    uint   `json:"schoolId"`
}

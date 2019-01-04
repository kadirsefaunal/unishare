package models

import (
	"github.com/jinzhu/gorm"
)

type School struct {
	gorm.Model
	Name    string  `json:"name" gorm:"size:255"`
	Classes []Class `gorm:"foreignkey:SchoolID"`
}

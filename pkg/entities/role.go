package entities

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
	Description string `json:"description"`
}
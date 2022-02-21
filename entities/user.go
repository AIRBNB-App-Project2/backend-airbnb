package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	User_uid string
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Rule     string `gorm:"type:enum('customer', 'owner');default:'customer'"`
}

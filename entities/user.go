package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	User_uid string `gorm:"unique;type:varchar(22)"`
	Name     string `gorm:"type:varchar(22)"`
	Email    string `gorm:"unique;index;not null;type:varchar(100)"`
	Status   string `gorm:"not null;type:varchar(100)"`
	Password string `gorm:"not null;type:varchar(100)"`
	// Room     []Room    `gorm:"foreignKey:User_uid;references:User_uid"`
	// Bookings []Booking `gorm:"foreignKey:User_uid;references:User_uid"`
	// Orders   []Order   `gorm:"foreignKey:User_uid;references:User_uid"`
}

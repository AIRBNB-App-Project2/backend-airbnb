package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	User_uid string
	Name     string
	Email    string    `gorm:"unique;index;not null;type:varchar(100)"`
	Password string    `gorm:"not null;type:varchar(100)"`
	Rule     string    `gorm:"type:enum('customer', 'owner');default:'customer'"`
	Room     []Room    `gorm:"foreignKey:User_uid;references:User_uid"`
	Bookings []Booking `gorm:"foreignKey:User_uid;references:User_uid"`
	Orders   []Order   `gorm:"foreignKey:User_uid;references:User_uid"`
}

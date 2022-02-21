package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Order_uid   string
	User_uid    string
	Booking_uid string
	Status      string `gorm:"type:enum('unpayed', 'payed');default:'unpayed'"`
}

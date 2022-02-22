package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Order_uid   string	`gorm:"unique;type:varchar(22)"`
	User_uid    string
	Booking_uid string
	Status      string `gorm:"type:enum('unpayed', 'payed');default:'unpayed'"`
}

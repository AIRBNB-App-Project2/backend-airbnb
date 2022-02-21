package entities

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	Booking_uid string
	User_uid    string
	Room_uid    string
	Start_date  datatypes.Date
	End_date    datatypes.Date
	Guest       int
	Status      string  `gorm:"type:enum('waiting', 'cancel', 'onGoing','end');default:'waiting'"`
	Orders      []Order `gorm:"foreignKey:Booking_uid;references:Booking_uid"`
}

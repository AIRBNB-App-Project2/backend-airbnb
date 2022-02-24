package entities

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	Booking_uid   string `gorm:"index;type:varchar(22)"`
	User_uid      string
	Room_uid      string
	Start_date    datatypes.Date
	End_date      datatypes.Date
	PaymentMethod string `gorm:"type:enum('qris');default:'qris'"`
	Status        string `gorm:"type:enum('waiting', 'cancel', 'reservation', 'onGoing','end');default:'waiting'"`
}

package entities

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Room_uid string	`gorm:"unique;type:varchar(22)"`
	User_uid string
	City_id  string
	Name     string
	Rating   int `gorm:"type:TINYINT"`
	Price    int
	Detail   string
	Images   []Image   `gorm:"foreignKey:Room_uid;references:Room_uid"`
	Bookings []Booking `gorm:"foreignKey:Room_uid;references:Room_uid"`
}

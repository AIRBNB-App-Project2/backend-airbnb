package entities

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Room_uid string
	User_uid string
	City_id  string
	Name     string
	Rating   int `gorm:"type:TINYINT"`
	Capacity int
	Price    int
	Detail   string
	Images   []Image   `gorm:"foreignKey:Room_uid;references:Room_uid"`
	Bookings []Booking `gorm:"foreignKey:Room_uid;references:Room_uid"`
}

package entities

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Room_uid string `gorm:"index;type:varchar(22)"`
	User_uid string 
	City_id  int
	Name     string
	Category string `gorm:"type:enum('classic', 'premium');default:'classic'"`
	Price    int
	Detail   string
	Images   []Image   `gorm:"foreignKey:Room_uid;references:Room_uid"`
	Bookings []Booking `gorm:"foreignKey:Room_uid;references:Room_uid"`
}

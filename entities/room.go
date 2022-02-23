package entities

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Room_uid    string `gorm:"index;type:varchar(22)"`
	User_uid    string
	City_id     int
	Address     string
	Name        string
	Category    string `gorm:"type:enum('standart', 'deluxe', 'superior', 'luxury');default:'standart'"`
	Status      string `gorm:"type:enum('open','close');default:'open'"`
	Price       int
	Description string
	Images      []Image   `gorm:"foreignKey:Room_uid;references:Room_uid"`
	Bookings    []Booking `gorm:"foreignKey:Room_uid;references:Room_uid"`
}

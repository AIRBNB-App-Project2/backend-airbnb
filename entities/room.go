package entities

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Room_uid string
	Name     string
	Rating   int `gorm:"type:TINYINT"`
	Capacity int
	Price    int
	Detail   string
}

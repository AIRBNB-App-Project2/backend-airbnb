package entities

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Room_uid string
	Url      string `gorm:"default:'https://www.teralogistics.com/wp-content/uploads/2020/12/default.png'"`
}

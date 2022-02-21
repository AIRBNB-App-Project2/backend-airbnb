package entities

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Image_uid string
	Room_uid  string
	Image     string `gorm:"default:'https://www.teralogistics.com/wp-content/uploads/2020/12/default.png'"`
}

package entities

import "gorm.io/gorm"

type Image struct {
	gorm.Config
	Image_uid string
	Image     string `gorm:"default:'https://www.teralogistics.com/wp-content/uploads/2020/12/default.png'"`
}

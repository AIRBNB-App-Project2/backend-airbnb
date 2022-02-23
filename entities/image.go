package entities

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Room_uid string
	Url      string `gorm:"default:'https://test-upload-s3-rogerdev.s3.ap-southeast-1.amazonaws.com/6216503718eb9324b8213a1f.png'"`
}

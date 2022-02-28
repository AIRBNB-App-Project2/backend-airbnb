package entities

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Room_uid string
	Url      string `gorm:"default:'https://karen-givi-bucket.s3.ap-southeast-1.amazonaws.com/621ce06818eb932118627489.png'"`
}

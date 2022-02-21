package entities

import "gorm.io/gorm"

type City struct {
	gorm.Model
	Province_id int
	Name        string
}

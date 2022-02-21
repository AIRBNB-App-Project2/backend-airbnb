package entities

import "gorm.io/gorm"

type Provinces struct {
	gorm.Model
	Name string
}

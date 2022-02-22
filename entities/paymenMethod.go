package entities

import "gorm.io/gorm"

type PaymentMethod struct {
	gorm.Model
	Name string `gorm:"unique;index;not null;type:varchar(100)"`
}
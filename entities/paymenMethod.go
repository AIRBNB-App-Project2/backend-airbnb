package entities

import "gorm.io/gorm"

type PaymentMethod struct {
	gorm.Model
	Name string
}

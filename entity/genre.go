package entity

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name string
}

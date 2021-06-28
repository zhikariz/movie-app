package entity

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserID  uint
	MovieID uint
	Review  string
	Rate    uint
	Movie   Movie
	User    User
}

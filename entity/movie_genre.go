package entity

import "gorm.io/gorm"

type MovieGenre struct {
	gorm.Model
	MovieID uint
	GenreID uint
	Movie   Movie
	Genre   Genre
}

package entity

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Ratings float64
	Title   string
	Year    time.Time
}

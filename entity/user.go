package entity

import "gorm.io/gorm"

type UserRole string

const (
	Admin = "Admin"
	Guest = "Guest"
)

type User struct {
	gorm.Model
	Email    string
	Password string
	Fullname string
	Role     UserRole
}

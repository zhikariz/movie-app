package user

import (
	. "movie-app/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]User, error)
	FindById(id int) (User, error)
	FindByEmail(email string) (User, error)
	Save(user User) (User, error)
	Update(user User) (User, error)
	Delete(user User) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *repository) FindById(id int) (User, error) {
	var user User
	err := r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Delete(user User) (bool, error) {
	err := r.db.Delete(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

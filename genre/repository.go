package genre

import (
	. "movie-app/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Genre, error)
	FindById(id int) (Genre, error)
	Save(genre Genre) (Genre, error)
	Update(genre Genre) (Genre, error)
	Delete(genre Genre) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Genre, error) {
	var genres []Genre
	err := r.db.Find(&genres).Error
	if err != nil {
		return genres, err
	}
	return genres, nil
}

func (r *repository) FindById(id int) (Genre, error) {
	var genre Genre
	err := r.db.Where("id = ?", id).Find(&genre).Error
	if err != nil {
		return genre, err
	}
	return genre, nil
}
func (r *repository) Save(genre Genre) (Genre, error) {
	err := r.db.Create(&genre).Error
	if err != nil {
		return genre, err
	}
	return genre, nil
}

func (r *repository) Update(genre Genre) (Genre, error) {
	err := r.db.Save(&genre).Error
	if err != nil {
		return genre, err
	}
	return genre, nil
}

func (r *repository) Delete(genre Genre) (bool, error) {
	err := r.db.Delete(&genre).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

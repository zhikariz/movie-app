package movie

import (
	. "movie-app/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Movie, error)
	FindById(id int) (Movie, error)
	Save(movie Movie) (Movie, error)
	Update(movie Movie) (Movie, error)
	Delete(movie Movie) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Movie, error) {
	var movies []Movie
	err := r.db.Find(&movies).Error
	if err != nil {
		return movies, err
	}
	return movies, nil

}

func (r *repository) FindById(id int) (Movie, error) {
	var movie Movie
	err := r.db.Where("id = ?", id).Find(&movie).Error
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func (r *repository) Save(movie Movie) (Movie, error) {
	err := r.db.Create(&movie).Error
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func (r *repository) Update(movie Movie) (Movie, error) {
	err := r.db.Save(&movie).Error
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func (r *repository) Delete(movie Movie) (bool, error) {
	err := r.db.Delete(&movie).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

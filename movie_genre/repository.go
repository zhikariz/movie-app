package movie_genre

import (
	. "movie-app/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]MovieGenre, error)
	FindById(id int) (MovieGenre, error)
	Save(movieGenre MovieGenre) (MovieGenre, error)
	Update(movieGenre MovieGenre) (MovieGenre, error)
	Delete(movieGenre MovieGenre) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewMovieGenreRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]MovieGenre, error) {
	var movieGenres []MovieGenre
	err := r.db.Preload("Movie").Preload("Genre").Find(&movieGenres).Error
	if err != nil {
		return movieGenres, err
	}
	return movieGenres, nil
}

func (r *repository) FindById(id int) (MovieGenre, error) {
	var movieGenre MovieGenre
	err := r.db.Where("id = ?", id).Preload("Movie").Preload("Genre").Find(&movieGenre).Error
	if err != nil {
		return movieGenre, err
	}
	return movieGenre, nil
}

func (r *repository) Save(movieGenre MovieGenre) (MovieGenre, error) {
	err := r.db.Create(&movieGenre).Error
	if err != nil {
		return movieGenre, err
	}
	newMovieGenre, _ := r.FindById(int(movieGenre.ID))
	return newMovieGenre, nil
}

func (r *repository) Update(movieGenre MovieGenre) (MovieGenre, error) {
	err := r.db.Save(&movieGenre).Error
	if err != nil {
		return movieGenre, err
	}
	return movieGenre, nil
}

func (r *repository) Delete(movieGenre MovieGenre) (bool, error) {
	err := r.db.Delete(&movieGenre).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

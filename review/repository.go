package review

import (
	. "movie-app/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Review, error)
	FindById(id int) (Review, error)
	Save(review Review) (Review, error)
	Update(review Review) (Review, error)
	Delete(review Review) (bool, error)
	FindByMovieId(movieId int) ([]Review, error)
}

type repository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Review, error) {
	var reviews []Review
	err := r.db.Preload("Movie").Preload("User").Find(&reviews).Error
	if err != nil {
		return reviews, err
	}
	return reviews, nil
}

func (r *repository) FindById(id int) (Review, error) {
	var review Review
	err := r.db.Preload("Movie").Preload("User").Where("id = ?", id).Error
	if err != nil {
		return review, err
	}
	return review, nil
}

func (r *repository) Save(review Review) (Review, error) {
	err := r.db.Create(&review).Error
	if err != nil {
		return review, err
	}
	newReview, err := r.FindById(int(review.ID))
	if err != nil {
		return newReview, err
	}
	return newReview, nil
}

func (r *repository) Update(review Review) (Review, error) {
	err := r.db.Save(&review).Error
	if err != nil {
		return review, err
	}
	return review, nil
}

func (r *repository) Delete(review Review) (bool, error) {
	err := r.db.Delete(&review).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repository) FindByMovieId(movieId int) ([]Review, error) {
	var reviews []Review
	err := r.db.Where("movie_id = ?", movieId).Preload("Movie").Preload("User").Find(&reviews).Error
	if err != nil {
		return reviews, err
	}
	return reviews, nil
}

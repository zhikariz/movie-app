package review

import (
	"errors"
	. "movie-app/entity"
	"strconv"
)

type Service interface {
	GetAll() ([]Review, error)
	GetById(uri GetReviewUriInput) (Review, error)
	Create(input CreateReviewInput) (Review, error)
	Update(uri GetReviewUriInput, input UpdateReviewInput) (Review, error)
	Delete(uri GetReviewUriInput) (bool, error)
	GetByMovieId(movieId int) ([]Review, error)
}

type service struct {
	repository Repository
}

func NewReviewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]Review, error) {
	reviews, err := s.repository.FindAll()
	if err != nil {
		return reviews, err
	}
	return reviews, nil
}

func (s *service) GetById(uri GetReviewUriInput) (Review, error) {
	review, err := s.repository.FindById(uri.ID)
	if err != nil {
		return review, err
	}
	return review, nil
}

func (s *service) Create(input CreateReviewInput) (Review, error) {
	movieId, _ := strconv.ParseUint(input.MovieID, 10, 32)
	rate, _ := strconv.ParseUint(input.Rate, 10, 32)
	review := Review{
		UserID:  input.UserID,
		MovieID: uint(movieId),
		Review:  input.Review,
		Rate:    uint(rate),
	}

	newReview, err := s.repository.Save(review)
	if err != nil {
		return newReview, err
	}
	return newReview, nil
}

func (s *service) Update(uri GetReviewUriInput, input UpdateReviewInput) (Review, error) {
	review, err := s.repository.FindById(uri.ID)
	if err != nil {
		return review, err
	}

	if review.ID == 0 {
		return review, errors.New("there is no reviews found with that id")
	}

	movieId, _ := strconv.ParseUint(input.MovieID, 10, 32)
	rate, _ := strconv.ParseUint(input.Rate, 10, 32)

	review.MovieID = uint(movieId)
	review.Rate = uint(rate)
	review.Review = input.Review

	updatedReview, err := s.repository.Update(review)
	if err != nil {
		return updatedReview, err
	}
	return updatedReview, nil
}

func (s *service) Delete(uri GetReviewUriInput) (bool, error) {
	review, err := s.repository.FindById(uri.ID)
	if err != nil {
		return false, err
	}

	if review.ID == 0 {
		return false, errors.New("there is no reviews found with that id")
	}

	isDeleted, err := s.repository.Delete(review)
	if err != nil {
		return isDeleted, err
	}
	return isDeleted, nil
}

func (s *service) GetByMovieId(movieId int) ([]Review, error) {
	reviews, err := s.repository.FindByMovieId(movieId)
	if err != nil {
		return reviews, err
	}
	return reviews, nil
}

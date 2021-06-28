package movie

import (
	"errors"
	. "movie-app/entity"
	"strconv"
	"time"
)

type Service interface {
	GetAll() ([]Movie, error)
	GetById(uri GetMovieUriInput) (Movie, error)
	Create(input CreateMovieInput) (Movie, error)
	Update(uri GetMovieUriInput, input UpdateMovieInput) (Movie, error)
	Delete(uri GetMovieUriInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewMovieService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]Movie, error) {
	movies, err := s.repository.FindAll()
	if err != nil {
		return movies, err
	}
	return movies, nil
}

func (s *service) GetById(uri GetMovieUriInput) (Movie, error) {
	movie, err := s.repository.FindById(uri.ID)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func (s *service) Create(input CreateMovieInput) (Movie, error) {
	ratings, err := strconv.ParseFloat(input.Ratings, 64)

	if err != nil {
		return Movie{}, err
	}

	year, err := time.Parse("2006", input.Year)

	if err != nil {
		return Movie{}, err
	}
	movie := Movie{
		Ratings: ratings,
		Title:   input.Title,
		Year:    year,
	}

	newMovie, err := s.repository.Save(movie)

	if err != nil {
		return newMovie, err
	}
	return newMovie, nil

}

func (s *service) Update(uri GetMovieUriInput, input UpdateMovieInput) (Movie, error) {
	movie, err := s.repository.FindById(uri.ID)

	if err != nil {
		return movie, err
	}

	if movie.ID == 0 {
		return movie, errors.New("there is no movie found with that id")
	}

	movie.Title = input.Title
	movie.Ratings, _ = strconv.ParseFloat(input.Ratings, 64)
	movie.Year, _ = time.Parse("2006", input.Year)

	updatedMovie, err := s.repository.Update(movie)

	if err != nil {
		return updatedMovie, err
	}

	return updatedMovie, nil
}

func (s *service) Delete(uri GetMovieUriInput) (bool, error) {
	movie, err := s.repository.FindById(uri.ID)

	if err != nil {
		return false, err
	}

	if movie.ID == 0 {
		return false, errors.New("there is no movie found with that id")
	}

	isUpdated, err := s.repository.Delete(movie)

	if err != nil {
		return isUpdated, err
	}

	return isUpdated, nil
}

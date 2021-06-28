package movie_genre

import (
	"errors"
	. "movie-app/entity"
	"strconv"
)

type Service interface {
	GetAll() ([]MovieGenre, error)
	GetById(uri GetUriMovieGenreInput) (MovieGenre, error)
	Create(input CreateMovieGenreInput) (MovieGenre, error)
	Update(uri GetUriMovieGenreInput, input UpdateMovieGenreInput) (MovieGenre, error)
	Delete(uri GetUriMovieGenreInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewMovieGenreService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]MovieGenre, error) {
	movieGenres, err := s.repository.FindAll()
	if err != nil {
		return movieGenres, err
	}
	return movieGenres, nil
}

func (s *service) GetById(uri GetUriMovieGenreInput) (MovieGenre, error) {
	movieGenre, err := s.repository.FindById(uri.ID)
	if err != nil {
		return movieGenre, err
	}
	return movieGenre, nil
}

func (s *service) Create(input CreateMovieGenreInput) (MovieGenre, error) {
	genreID, _ := strconv.ParseUint(input.GenreID, 10, 32)
	movieID, _ := strconv.ParseUint(input.MovieID, 10, 32)

	movieGenre := MovieGenre{
		MovieID: uint(genreID),
		GenreID: uint(movieID),
	}

	newMovieGenre, err := s.repository.Save(movieGenre)
	if err != nil {
		return newMovieGenre, err
	}
	return newMovieGenre, nil
}

func (s *service) Update(uri GetUriMovieGenreInput, input UpdateMovieGenreInput) (MovieGenre, error) {
	movieGenre, err := s.repository.FindById(uri.ID)
	if err != nil {
		return movieGenre, err
	}

	if movieGenre.ID == 0 {
		return movieGenre, errors.New("there is no movie genre with that id")
	}

	genreID, _ := strconv.ParseUint(input.GenreID, 10, 32)
	movieID, _ := strconv.ParseUint(input.MovieID, 10, 32)
	movieGenre.GenreID = uint(genreID)
	movieGenre.MovieID = uint(movieID)

	updatedMovieGenre, err := s.repository.Update(movieGenre)
	if err != nil {
		return updatedMovieGenre, err
	}
	return updatedMovieGenre, nil
}

func (s *service) Delete(uri GetUriMovieGenreInput) (bool, error) {
	movie, err := s.repository.FindById(uri.ID)
	if err != nil {
		return false, err
	}
	if movie.ID == 0 {
		return false, errors.New("there is no movie genre with that id")
	}

	isDeleted, err := s.repository.Delete(movie)

	if err != nil {
		return isDeleted, err
	}
	return isDeleted, nil
}

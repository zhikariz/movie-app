package genre

import (
	"errors"
	. "movie-app/entity"
)

type Service interface {
	GetAll() ([]Genre, error)
	GetById(uri GetGenreUriInput) (Genre, error)
	Create(input CreateGenreInput) (Genre, error)
	Update(uri GetGenreUriInput, input UpdateGenreInput) (Genre, error)
	//Delete(uri GetGenreUriInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewGenreService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]Genre, error) {
	genres, err := s.repository.FindAll()
	if err != nil {
		return genres, err
	}
	return genres, nil
}

func (s *service) GetById(uri GetGenreUriInput) (Genre, error) {
	genre, err := s.repository.FindById(uri.ID)
	if err != nil {
		return genre, err
	}
	return genre, nil
}

func (s *service) Create(input CreateGenreInput) (Genre, error) {
	genre := Genre{
		Name: input.Name,
	}
	genre, err := s.repository.Save(genre)
	if err != nil {
		return genre, err
	}
	return genre, nil
}

func (s *service) Update(uri GetGenreUriInput, input UpdateGenreInput) (Genre, error) {
	genre, err := s.repository.FindById(uri.ID)
	if err != nil {
		return genre, err
	}

	if genre.ID == 0 {
		return genre, errors.New("there is no genre with that id")
	}

	genre.Name = input.Name

	updatedGenre, err := s.repository.Update(genre)

	if err != nil {
		return updatedGenre, err
	}

	return updatedGenre, nil
}

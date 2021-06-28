package user

import (
	"errors"
	. "movie-app/entity"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterUserInput) (User, error)
	Login(input LoginUserInput) (User, error)
	GetUserById(input GetUserUriInput) (User, error)
}

type service struct {
	repository Repository
}

func NewUserService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(input RegisterUserInput) (User, error) {
	user := User{
		Email:    input.Email,
		Fullname: input.Fullname,
		Role:     Guest,
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	checkUser, err := s.repository.FindByEmail(user.Email)

	if err != nil {
		return checkUser, err
	}

	if checkUser.ID != 0 {
		return checkUser, errors.New("email is exist at the database")
	}

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginUserInput) (User, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return user, err
	}
	return user, nil

}

func (s *service) GetUserById(input GetUserUriInput) (User, error) {
	user, err := s.repository.FindById(input.ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("there is no user with that id")
	}

	return user, nil
}

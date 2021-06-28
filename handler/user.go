package handler

import (
	"movie-app/auth"
	"movie-app/helper"
	"movie-app/user"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userHandler struct {
	authService auth.Service
	userService user.Service
}

func NewUserHandler(authService auth.Service, userService user.Service) *userHandler {
	return &userHandler{authService, userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(c, err, "Register User Error")
		return
	}

	newUser, err := h.userService.Register(input)

	if err != nil {
		helper.ErrorHandling(c, err, "Register User Error")
		return
	}

	formatter := user.FormatUser(newUser)
	helper.SuccessHandling(c, formatter, "User Successfully Registered")
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(c, err, "Login Failed !")
		return
	}

	loggedInUser, err := h.userService.Login(input)
	if err != nil {
		helper.ErrorHandling(c, err, "Login Failed !")
		return
	}

	token, err := h.authService.GenerateToken(loggedInUser)
	if err != nil {
		helper.ErrorHandling(c, err, "Login Failed !")
		return
	}

	formatter := user.FormatUserToken(loggedInUser, token)
	helper.SuccessHandling(c, formatter, "Successfully Logged in !")

}

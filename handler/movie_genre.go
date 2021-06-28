package handler

import (
	"movie-app/helper"
	"movie-app/movie_genre"

	"github.com/gin-gonic/gin"
)

type MovieGenreHandler interface {
	GetAllMovieGenre(c *gin.Context)
	CreateMovieGenre(c *gin.Context)
}

type movieGenreHandler struct {
	movieGenreService movie_genre.Service
}

func NewMovieGenreHandler(movieGenreService movie_genre.Service) *movieGenreHandler {
	return &movieGenreHandler{movieGenreService}
}

func (h *movieGenreHandler) GetAllMovieGenre(c *gin.Context) {
	movieGenres, err := h.movieGenreService.GetAll()
	if err != nil {
		helper.ErrorHandling(c, err, "Failed to showing all movie genres")
		return
	}
	formatter := movie_genre.FormatMovieGenres(movieGenres)
	helper.SuccessHandling(c, formatter, "Successfully showing all movie genres")
}

func (h *movieGenreHandler) CreateMovieGenre(c *gin.Context) {
	var input movie_genre.CreateMovieGenreInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(c, err, "Failed to create movie genre")
		return
	}

	movieGenre, err := h.movieGenreService.Create(input)

	if err != nil {
		helper.ErrorHandling(c, err, "Failed to create movie genre")
		return
	}

	formatter := movie_genre.FormatMovieGenre(movieGenre)
	helper.SuccessHandling(c, formatter, "Successfully create movie genre")
}

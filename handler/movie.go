package handler

import (
	"movie-app/helper"
	"movie-app/movie"

	"github.com/gin-gonic/gin"
)

type MovieHandler interface {
	GetAllMovie(c *gin.Context)
	CreateMovie(c *gin.Context)
}

type movieHandler struct {
	movieService movie.Service
}

func NewMovieHandler(movieService movie.Service) *movieHandler {
	return &movieHandler{movieService}
}

func (h *movieHandler) GetAllMovie(c *gin.Context) {
	movies, err := h.movieService.GetAll()
	if err != nil {
		helper.ErrorHandling(c, err, "Failed to Showing All Movies")
		return
	}

	formatter := movie.FormatMovies(movies)
	helper.SuccessHandling(c, formatter, "Successfully showing all movies")
}

func (h *movieHandler) CreateMovie(c *gin.Context) {
	var input movie.CreateMovieInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(c, err, "Failed to Create a Movie")
		return
	}

	newMovie, err := h.movieService.Create(input)

	if err != nil {
		helper.ErrorHandling(c, err, "Failed to Create a Movie")
		return
	}

	formatter := movie.FormatMovie(newMovie)
	helper.SuccessHandling(c, formatter, "Successfully create a movie")
}

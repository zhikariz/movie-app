package handler

import (
	"movie-app/genre"
	"movie-app/helper"

	"github.com/gin-gonic/gin"
)

type GenreHandler interface {
	CreateGenre(c *gin.Context)
	GetAllGenre(c *gin.Context)
	UpdateGenre(c *gin.Context)
}

type genreHandler struct {
	genreService genre.Service
}

func NewGenreHandler(genreService genre.Service) *genreHandler {
	return &genreHandler{genreService}
}

func (h *genreHandler) CreateGenre(c *gin.Context) {
	var input genre.CreateGenreInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(c, err, "Create Genre Error")
		return
	}

	newGenre, err := h.genreService.Create(input)

	if err != nil {
		helper.ErrorHandling(c, err, "Create Genre Error")
		return
	}

	formatter := genre.FormatGenre(newGenre)
	helper.SuccessHandling(c, formatter, "Successfully Create Genre")
}

func (h *genreHandler) GetAllGenre(c *gin.Context) {
	genres, err := h.genreService.GetAll()

	if err != nil {
		helper.ErrorHandling(c, err, "Failed to Showing All Genres")
		return
	}

	formatter := genre.FormatGenres(genres)
	helper.SuccessHandling(c, formatter, "Successfully Showing All Genres")
}

func (h *genreHandler) UpdateGenre(c *gin.Context) {
	var uri genre.GetGenreUriInput
	var input genre.UpdateGenreInput

	err := c.ShouldBindUri(&uri)
	if err != nil {
		helper.ErrorValidation(c, err, "Failed to Update Genre")
		return
	}

	err = c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(c, err, "Failed to Update Genre")
		return
	}

	updatedGenre, err := h.genreService.Update(uri, input)

	if err != nil {
		helper.ErrorHandling(c, err, "Failed to Update Genre")
		return
	}

	formatter := genre.FormatGenre(updatedGenre)
	helper.SuccessHandling(c, formatter, "Successfully update Genre")
}

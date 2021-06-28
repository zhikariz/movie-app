package handler

import (
	. "movie-app/entity"
	"movie-app/helper"
	"movie-app/review"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewHandler interface {
	GetAllReview(c *gin.Context)
	CreateReview(c *gin.Context)
	GetReviewByMovieId(c *gin.Context)
}

type reviewHandler struct {
	reviewService review.Service
}

func NewReviewHandler(reviewService review.Service) *reviewHandler {
	return &reviewHandler{reviewService}
}

func (h *reviewHandler) GetAllReview(c *gin.Context) {
	reviews, err := h.reviewService.GetAll()
	if err != nil {
		helper.ErrorHandling(c, err, "Failed to showing all reviews")
		return
	}
	formatter := review.FormatReviews(reviews)
	helper.SuccessHandling(c, formatter, "Successfully to showing all reviews")

}

func (h *reviewHandler) CreateReview(c *gin.Context) {
	var input review.CreateReviewInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(c, err, "Failed to Create Review")
		return
	}

	currentUser := c.MustGet("currentUser").(User)

	input.UserID = uint(currentUser.ID)

	newReview, err := h.reviewService.Create(input)

	if err != nil {
		helper.ErrorHandling(c, err, "Failed to Create Review")
		return
	}

	formatter := review.FormatReview(newReview)
	helper.SuccessHandling(c, formatter, "Successfully Create Review")
}

func (h *reviewHandler) GetReviewByMovieId(c *gin.Context) {
	movieID := c.Query("movieId")
	movieId, err := strconv.Atoi(movieID)

	if err != nil {
		helper.ErrorHandling(c, err, "Failed to Get Reviews By Movie ID")
		return
	}

	reviews, err := h.reviewService.GetByMovieId(movieId)

	if err != nil {
		helper.ErrorHandling(c, err, "Failed to Get Reviews By Movie ID")
		return
	}
	formatter := review.FormatReviews(reviews)
	helper.SuccessHandling(c, formatter, "Sucessfully showing reviews by movies id")

}

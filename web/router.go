package web

import (
	"movie-app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouterV1(api *gin.RouterGroup, db *gorm.DB) {

	repository := GetRepository(db)
	service := GetService(repository)
	module := GetModule(service)

	api.POST("/register", module.user.RegisterUser)
	api.POST("/login", module.user.LoginUser)

	api.POST("/genres", middleware.AdminMiddleware(service.auth, service.user), module.genre.CreateGenre)
	api.GET("/genres", module.genre.GetAllGenre)
	api.PUT("/genres/:id", middleware.AdminMiddleware(service.auth, service.user), module.genre.UpdateGenre)

	api.GET("/movies", module.movie.GetAllMovie)
	api.POST("/movies", middleware.AdminMiddleware(service.auth, service.user), module.movie.CreateMovie)

	api.GET("/movie_genres", module.movieGenre.GetAllMovieGenre)
	api.POST("/movie_genres", middleware.AdminMiddleware(service.auth, service.user), module.movieGenre.CreateMovieGenre)

	api.GET("/reviews", module.review.GetAllReview)
	api.POST("/reviews", middleware.UserMiddleware(service.auth, service.user), module.review.CreateReview)
	api.GET("/review", module.review.GetReviewByMovieId)
}

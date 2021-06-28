package web

import (
	"movie-app/genre"
	"movie-app/movie"
	"movie-app/movie_genre"
	"movie-app/review"
	"movie-app/user"

	"gorm.io/gorm"
)

type handlerRepository struct {
	user       user.Repository
	genre      genre.Repository
	movie      movie.Repository
	movieGenre movie_genre.Repository
	review     review.Repository
}

func GetRepository(db *gorm.DB) handlerRepository {
	userRepository := user.NewUserRepository(db)
	genreRepository := genre.NewGenreRepository(db)
	movieRepository := movie.NewMovieRepository(db)
	movieGenreRepository := movie_genre.NewMovieGenreRepository(db)
	reviewRepository := review.NewReviewRepository(db)

	return handlerRepository{
		user:       userRepository,
		genre:      genreRepository,
		movie:      movieRepository,
		movieGenre: movieGenreRepository,
		review:     reviewRepository,
	}
}

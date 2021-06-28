package web

import (
	"movie-app/auth"
	"movie-app/genre"
	"movie-app/movie"
	"movie-app/movie_genre"
	"movie-app/review"
	"movie-app/user"
)

type handlerService struct {
	auth       auth.Service
	user       user.Service
	genre      genre.Service
	movie      movie.Service
	movieGenre movie_genre.Service
	review     review.Service
}

func GetService(repo handlerRepository) handlerService {
	authService := auth.NewService()
	userService := user.NewUserService(repo.user)
	genreService := genre.NewGenreService(repo.genre)
	movieService := movie.NewMovieService(repo.movie)
	movieGenreService := movie_genre.NewMovieGenreService(repo.movieGenre)
	reviewService := review.NewReviewService(repo.review)

	return handlerService{
		auth:       authService,
		user:       userService,
		genre:      genreService,
		movie:      movieService,
		movieGenre: movieGenreService,
		review:     reviewService,
	}
}

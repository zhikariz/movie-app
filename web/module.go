package web

import (
	"movie-app/handler"
)

type handlerModule struct {
	user       handler.UserHandler
	genre      handler.GenreHandler
	movie      handler.MovieHandler
	movieGenre handler.MovieGenreHandler
	review     handler.ReviewHandler
}

func GetModule(service handlerService) handlerModule {
	userHandler := handler.NewUserHandler(service.auth, service.user)
	genreHandler := handler.NewGenreHandler(service.genre)
	movieHandler := handler.NewMovieHandler(service.movie)
	movieGenreHandler := handler.NewMovieGenreHandler(service.movieGenre)
	reviewHandler := handler.NewReviewHandler(service.review)

	return handlerModule{
		user:       userHandler,
		genre:      genreHandler,
		movie:      movieHandler,
		movieGenre: movieGenreHandler,
		review:     reviewHandler,
	}
}

package movie_genre

import (
	. "movie-app/entity"
	"movie-app/genre"
	"movie-app/movie"
)

type MovieGenreFormatter struct {
	ID      int                  `json:"id"`
	MovieID uint                 `json:"movie_id"`
	GenreID uint                 `json:"genre_id"`
	Movie   movie.MovieFormatter `json:"movie"`
	Genre   genre.GenreFormatter `json:"genre"`
}

func FormatMovieGenre(movieGenre MovieGenre) (movieGenreFormatter MovieGenreFormatter) {
	movieGenreFormatter = MovieGenreFormatter{
		ID:      int(movieGenre.ID),
		MovieID: movieGenre.MovieID,
		GenreID: movieGenre.GenreID,
		Movie:   movie.FormatMovie(movieGenre.Movie),
		Genre:   genre.FormatGenre(movieGenre.Genre),
	}
	return
}

func FormatMovieGenres(movieGenres []MovieGenre) (movieGenreFormatters []MovieGenreFormatter) {
	if len(movieGenres) == 0 {
		return []MovieGenreFormatter{}
	}

	for _, movieGenre := range movieGenres {
		formatter := FormatMovieGenre(movieGenre)
		movieGenreFormatters = append(movieGenreFormatters, formatter)
	}
	return
}

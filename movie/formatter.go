package movie

import (
	. "movie-app/entity"
)

type MovieFormatter struct {
	ID      int     `json:"id"`
	Ratings float64 `json:"ratings"`
	Title   string  `json:"title"`
	Year    int     `json:"year"`
}

func FormatMovie(movie Movie) (movieFormatter MovieFormatter) {
	movieFormatter = MovieFormatter{
		ID:      int(movie.ID),
		Ratings: movie.Ratings,
		Title:   movie.Title,
		Year:    movie.Year.Year(),
	}
	return
}

func FormatMovies(movies []Movie) (movieFormatters []MovieFormatter) {
	if len(movies) == 0 {
		return []MovieFormatter{}
	}

	for _, movie := range movies {
		formatter := FormatMovie(movie)
		movieFormatters = append(movieFormatters, formatter)
	}
	return
}

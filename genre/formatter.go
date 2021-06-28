package genre

import . "movie-app/entity"

type GenreFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatGenre(genre Genre) (genreFormatter GenreFormatter) {
	genreFormatter = GenreFormatter{
		ID:   int(genre.ID),
		Name: genre.Name,
	}
	return
}

func FormatGenres(genres []Genre) (genreFormatters []GenreFormatter) {
	if len(genres) == 0 {
		return []GenreFormatter{}
	}

	for _, genre := range genres {
		formatter := FormatGenre(genre)
		genreFormatters = append(genreFormatters, formatter)
	}
	return
}

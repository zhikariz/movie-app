package movie_genre

type CreateMovieGenreInput struct {
	MovieID string `json:"movie_id" binding:"required"`
	GenreID string `json:"genre_id" binding:"required"`
}

type UpdateMovieGenreInput struct {
	MovieID string `json:"movie_id" binding:"required"`
	GenreID string `json:"genre_id" binding:"required"`
}

type GetUriMovieGenreInput struct {
	ID int `uri:"id" binding:"required"`
}

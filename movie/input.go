package movie

type CreateMovieInput struct {
	Ratings string `json:"ratings" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Year    string `json:"year" binding:"required"`
}

type UpdateMovieInput struct {
	Ratings string `json:"ratings" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Year    string `json:"year" binding:"required"`
}

type GetMovieUriInput struct {
	ID int `uri:"id" binding:"required"`
}

package genre

type CreateGenreInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateGenreInput struct {
	Name string `json:"name" binding:"required"`
}

type GetGenreUriInput struct {
	ID int `uri:"id" binding:"required"`
}

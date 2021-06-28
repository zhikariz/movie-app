package review

type CreateReviewInput struct {
	UserID  uint
	MovieID string `json:"movie_id" binding:"required"`
	Review  string `json:"review" binding:"required"`
	Rate    string `json:"rate" binding:"required"`
}

type UpdateReviewInput struct {
	MovieID string `json:"movie_id" binding:"required"`
	Review  string `json:"review" binding:"required"`
	Rate    string `json:"rate" binding:"required"`
}

type GetReviewUriInput struct {
	ID int `uri:"id" binding:"required"`
}

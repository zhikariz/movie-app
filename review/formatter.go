package review

import (
	. "movie-app/entity"
	"movie-app/movie"
	"movie-app/user"
)

type ReviewFormatter struct {
	ID      int                  `json:"id"`
	UserID  uint                 `json:"user_id"`
	MovieID uint                 `json:"movie_id"`
	Review  string               `json:"review"`
	Rate    uint                 `json:"rate"`
	Movie   movie.MovieFormatter `json:"movie"`
	User    user.UserFormatter   `json:"user"`
}

func FormatReview(review Review) (reviewFormatter ReviewFormatter) {
	reviewFormatter = ReviewFormatter{
		ID:      int(review.ID),
		UserID:  review.UserID,
		MovieID: review.MovieID,
		Review:  review.Review,
		Rate:    review.Rate,
		Movie:   movie.FormatMovie(review.Movie),
		User:    user.FormatUser(review.User),
	}
	return
}

func FormatReviews(reviews []Review) (reviewFormatters []ReviewFormatter) {
	if len(reviews) == 0 {
		return []ReviewFormatter{}
	}

	for _, review := range reviews {
		formatter := FormatReview(review)
		reviewFormatters = append(reviewFormatters, formatter)
	}
	return
}

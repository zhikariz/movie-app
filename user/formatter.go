package user

import (
	. "movie-app/entity"
)

type UserFormatterToken struct {
	ID       uint     `json:"id"`
	Email    string   `json:"email"`
	Fullname string   `json:"full_name"`
	Role     UserRole `json:"role"`
	Token    string   `json:"token"`
}

type UserFormatter struct {
	ID       uint     `json:"id"`
	Email    string   `json:"email"`
	Fullname string   `json:"full_name"`
	Role     UserRole `json:"role"`
}

func FormatUserToken(user User, token string) (userFormatterToken UserFormatterToken) {
	userFormatterToken = UserFormatterToken{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Role:     user.Role,
		Token:    token,
	}
	return
}

func FormatUser(user User) (userFormatter UserFormatter) {
	userFormatter = UserFormatter{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Role:     user.Role,
	}
	return
}

func FormatUsers(users []User) (userFormatters []UserFormatter) {
	if len(users) == 0 {
		return []UserFormatter{}
	}

	for _, user := range users {
		formatter := FormatUser(user)
		userFormatters = append(userFormatters, formatter)
	}
	return
}

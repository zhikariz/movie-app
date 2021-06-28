package middleware

import (
	"movie-app/auth"
	"movie-app/helper"
	"movie-app/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			helper.AuthorizationHandling(c)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			helper.AuthorizationHandling(c)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			helper.AuthorizationHandling(c)
			return
		}

		userId := int(claim["user_id"].(float64))

		account, err := userService.GetUserById(user.GetUserUriInput{ID: userId})

		if err != nil {
			helper.AuthorizationHandling(c)
			return
		}

		if claim["role"] != "Guest" {
			helper.AuthorizationHandling(c)
			return
		}

		c.Set("currentUser", account)
	}
}

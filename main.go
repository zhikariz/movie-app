package main

import (
	"movie-app/helper"
	"movie-app/web"

	"github.com/gin-gonic/gin"
)

func main() {
	helper.Load()

	router := gin.Default()
	db := helper.SetupDB()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	api := router.Group("api/v1")
	web.SetupRouterV1(api, db)

	router.Run()
}

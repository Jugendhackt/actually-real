package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main/app"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/me/friends/list", func(c *gin.Context) {
		c.String(http.StatusOK, "Not implemented yet.")
	})

	r.GET("/me/friends/requests/self", func(c *gin.Context) {
		c.String(http.StatusOK, "Not implemented yet.")
	})

	r.POST("/me/friends/add", func(c *gin.Context) {

	})

	r.POST("/me/friends/requests/sent", func(c *gin.Context) {

	})

	r.POST("/me/img/upload", func(c *gin.Context) {

	})

	r.POST("/user/create", func(c *gin.Context) {
		var newUser app.User

		if err := c.BindJSON(&newUser); err != nil {
			return
		}

		c.IndentedJSON(http.StatusCreated, newUser)
	})

	r.POST("/user/login", func(c *gin.Context) {

	})

	return r
}

func StartApi() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

package handler

import (
	"net/http"

	"main/app"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Name string
}

func setupRouter(a *app.App) *gin.Engine {
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

	r.GET("/img/", func(c *gin.Context) {

	})

	r.POST("/me/friends/add", func(c *gin.Context) {

	})

	r.POST("/me/friends/requests/sent", func(c *gin.Context) {

	})

	r.POST("/me/img/upload", func(c *gin.Context) {

	})

	r.POST("/user/create", func(c *gin.Context) {
		req := createUserRequest{}

		if err := c.BindJSON(&req); err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		app.CreateUser(a, req.Name)

		c.Status(http.StatusCreated)
	})

	r.POST("/user/login", func(c *gin.Context) {

	})

	return r
}

func StartApi(app *app.App) {
	r := setupRouter(app)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"main/app"
	"main/database"

	"github.com/gin-gonic/gin"
)

type SendFriendRequest struct {
	Name   string
	Friend string
}

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
		entries, err := os.ReadDir("./images")
		if err != nil {
			log.Fatal(err)
		}

		resp := struct {
			Files []string
		}{
			[]string{},
		}

		for _, e := range entries {
			log.Println(e.Name())
			resp.Files = append(resp.Files, e.Name())
		}

		c.JSON(http.StatusOK, resp)
	})

	r.POST("/me/friends/add", func(c *gin.Context) {
		var req SendFriendRequest

		if err := c.BindJSON(&req); err != nil {
			return
		}
		var user database.User
		var friend database.User

		a.DB.Where("name = ?", req.Name).First(&user)
		a.DB.Where("name = ?", req.Friend).First(&friend)

		user.Friends = append(user.Friends, friend)
		friend.Friends = append(friend.Friends, user)

		fmt.Println(user)
		fmt.Println(friend)
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

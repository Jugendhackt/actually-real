package handler

import (
<<<<<<< HEAD
<<<<<<< HEAD
	"fmt"
	"log"
=======
>>>>>>> ac427007982396e147c34eed9ec50ff55dd2acd9
=======
	"fmt"
	"log"
>>>>>>> 47accf8e12c12c04a811da5631f4df1f20eebb97
	"net/http"
	"os"

	"main/app"

	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
type SendFriendRequest struct {
	Name   string
	Friend string
}

func setupRouter() *gin.Engine {
=======
type createUserRequest struct {
	Name string
}

func setupRouter(a *app.App) *gin.Engine {
>>>>>>> ac427007982396e147c34eed9ec50ff55dd2acd9
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
		var user app.User
		var friend app.User

		db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

		if err != nil {
			panic("failed to connect database")
		}

		db.Where("name = ?", req.Name).First(&user)
		db.Where("name = ?", req.Friend).First(&friend)

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

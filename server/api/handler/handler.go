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

type GetNameNameFriend struct {
	Name   string
	Friend string
}

type GetName struct {
	Name string
}

type GetPath struct {
	Path string
}

func setupRouter(a *app.App) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/me/img/list", func(c *gin.Context) {
		var req GetName

		if err := c.BindJSON(&req); err != nil {
			return
		}

		var user database.User = getUserFromDB(req.Name, a)

		var images []database.Image

		a.DB.Model(&user).Association("Images").Find(&images)

		c.JSON(http.StatusOK, images)
	})

	r.POST("/me/friends/list", func(c *gin.Context) {

		var req GetName

		if err := c.BindJSON(&req); err != nil {
			return
		}

		var user database.User = getUserFromDB(req.Name, a)

		var FriendList []database.User

		a.DB.Model(&user).Association("Friends").Find(&FriendList)

		c.JSON(http.StatusOK, FriendList)
	})

	r.GET("/me/friends/requests/self", func(c *gin.Context) {
		c.String(http.StatusOK, "Not implemented yet.")
	})

	r.POST("/img/", func(c *gin.Context) {
		req := GetPath{}

		if err := c.BindJSON(&req); err != nil {
			return
		}

		path := "images/" + req.Path

		_, err := os.Stat(path)
		if err == nil {
			c.File(path)
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	r.POST("/me/friends/add", func(c *gin.Context) {
		var req GetNameNameFriend

		if err := c.BindJSON(&req); err != nil {
			return
		}

		var user database.User
		var friend database.User

		a.DB.Where("name = ?", req.Name).First(&user)
		a.DB.Where("name = ?", req.Friend).First(&friend)

		user.Friends = append(user.Friends, friend)
		friend.Friends = append(friend.Friends, user)

		a.DB.Save(&user)
		a.DB.Save(&friend)

	})

	r.POST("/me/friends/requests/sent", func(c *gin.Context) {

	})

	r.POST("/me/img/upload", func(c *gin.Context) {
		name := c.PostForm("name")

		file, err := c.FormFile("image")
		if err != nil {
			log.Fatal(err)
		}

		err = c.SaveUploadedFile(file, "images/"+file.Filename)
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

		user := database.User{}
		a.DB.Where("name = ?", name).First(&user)

		image := database.Image{
			Path: file.Filename,
		}

		a.DB.Model(&user).Association("Images").Append(&image)
		a.DB.Save(&user)
	})

	r.POST("/user/create", func(c *gin.Context) {
		req := GetName{}

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

func getUserFromDB(name string, app *app.App) database.User {
	var user database.User
	app.DB.Where("name = ?", name).First(&user)
	return user
}

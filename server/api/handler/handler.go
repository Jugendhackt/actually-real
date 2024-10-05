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

func setupRouter(a *app.App) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/me/img/list", func(c *gin.Context) {
		req := GetName{}
		user := database.User{}

		if err := c.BindJSON(&req); err != nil {
			return
		}

		a.DB.Where("name = ?", req.Name).First(&user)
		log.Println(user)

		var images []database.Image
		a.DB.Model(&user).Association("Images").Find(&images)

		c.JSON(http.StatusOK, images)
	})

	r.POST("/me/friends/list", func(c *gin.Context) {
		var req GetName

		if err := c.BindJSON(&req); err != nil {
			return
		}

		var user database.User

		a.DB.Where("name = ?", req.Name).First(&user)

		var FriendList []database.User

		a.DB.Model(&user).Association("Friends").Find(&FriendList)

		fmt.Println(FriendList)

		c.JSON(http.StatusOK, FriendList)
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
	// WIP
	// r.POST("/me/friends/remove", func(c *gin.Context) {
	//   var req GetNameNameFriend
	//
	//   if err:= c.BindJSON(&req); err !=nil {
	//     return
	//   }
	//   var user database.User
	//   var friend database.User
	//
	// }
	//
	r.POST("/me/img/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		log.Println(name)
		file, err := c.FormFile("image")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(file.Filename)

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

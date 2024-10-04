package handler

import (
	"log"
	"net/http"

	"main/app"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
		db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

		if err != nil {
			panic("failed to connect database")
		}

		var newUser app.User

		if err := c.BindJSON(&newUser); err != nil {
			log.Println(newUser)
			return
		}

		db.Create(&newUser)

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

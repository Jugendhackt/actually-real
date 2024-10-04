package main

import (
	"main/api/handler"
	"main/app"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Create sqlite database
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Imports schema into database
	db.AutoMigrate(&app.User{})
	db.AutoMigrate(&app.Image{})

	handler.StartApi()
}

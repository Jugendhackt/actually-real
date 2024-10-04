package main

import (
	"main/api/handler"
	"main/app"
	"main/database"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {

	// Create sqlite database
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&database.User{})
	db.AutoMigrate(&database.Image{})

	return db
}

func initApp() app.App {
	db := initDB()

	return app.App{
		DB: db,
	}
}

func main() {
	app := initApp()

	handler.StartApi(&app)
}

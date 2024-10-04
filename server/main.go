package main

import (
	handler "main/api/handler"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID      uint `gorm:"primaryKey,autoIncrement"`
	Name    string
	Friends []User `gorm:"many2many:friends"`
	Images  []Image
}

type Image struct {
	ID      uint `gorm:"primaryKey,autoIncrement"`
	Path    string
	UserID  uint
	Created time.Time
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Image{})

	handler.Ping()
}

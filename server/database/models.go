package database

import "time"

type User struct {
	ID      uint   `gorm:"primaryKey,autoIncrement"`
	Name    string `gorm:"unique,not null"`
	Friends []User `gorm:"many2many:friends"`
	Images  []Image
}

type Image struct {
	ID      uint `gorm:"primaryKey,autoIncrement"`
	Path    string
	UserID  uint
	Created time.Time
}

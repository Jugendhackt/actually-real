package app

import (
	"time"
)

type Image struct {
	ID      uint `gorm:"primaryKey,autoIncrement"`
	Path    string
	UserID  uint
	Created time.Time
}

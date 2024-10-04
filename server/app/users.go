package app

import (
	"main/database"
)

func CreateUser(a *App, name string) {
	var user database.User

	if err := a.DB.Where("name = ?", name).First(&user).Error; err != nil {
		newUser := database.User{Name: name}
		a.DB.Create(&newUser)
		return
	}
}

package app

type User struct {
	ID      uint `gorm:"primaryKey,autoIncrement"`
	Name    string
	Friends []User `gorm:"many2many:friends"`
	Images  []Image
}

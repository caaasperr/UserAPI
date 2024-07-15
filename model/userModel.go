package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"size:255;unique"`
	Password string
}

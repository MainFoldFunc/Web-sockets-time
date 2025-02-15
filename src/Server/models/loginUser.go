package models

type UserLogin struct {
	Email    string `gorm:"unique"`
	Password string `gorm:"size:255"`
}

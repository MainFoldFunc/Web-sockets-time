package models

type Conv struct {
	Message uint32 `gorm:"primaryKey"`
	User1   string `gorm:"size:255"`
	User2   string `gorm:"size:255"`
	Body    string `gorm:"size:255"`
}

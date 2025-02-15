package models

type ChatReqest struct {
	ID     uint   `gorm:"primaryKey"`
	UserS  string `gorm:"size:255"`
	UserR  string `gorm:"size:255"`
	Status string `gorm:"size:255"`
}

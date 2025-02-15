package models

type SaveToDatabase struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"size:255"`
	Sender  int    // 1 for sender, 0 for receiver
}

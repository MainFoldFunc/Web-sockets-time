package models

type SeeChatReqests struct {
	UserEmail string `json:"userEmail" gorm:"size:255"` // Email of the user you're checking requests for
}

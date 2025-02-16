package models

// SearchForUsersBar is used to capture the search query sent from the frontend.
type SearchForUsersBar struct {
	Email string `json:"email"` // User's email to search for
}

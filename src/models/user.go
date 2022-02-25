package models

// User struct for CRUD
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

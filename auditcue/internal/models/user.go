package models

type User struct {
    ID       string `json:"id"`
    Email    string `json:"email"`
    Password string `json:"password"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}

// NewUser creates a new user instance
func NewUser(id, email, password, createdAt, updatedAt string) *User {
    return &User{
        ID:       id,
        Email:    email,
        Password: password,
        CreatedAt: createdAt,
        UpdatedAt: updatedAt,
    }
}
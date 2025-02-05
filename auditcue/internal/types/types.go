package types

// Request and response models for the AuditCue application

// SignupRequest represents the payload for user signup requests.
type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// SignupResponse represents the response returned after a successful signup.
type SignupResponse struct {
	Message string `json:"message"`
	UserID  string `json:"user_id"`
}

// LoginRequest represents the payload for user login requests.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the response returned after a successful login.
type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// ErrorResponse represents a standard error response format.
type ErrorResponse struct {
	Error string `json:"error"`
}
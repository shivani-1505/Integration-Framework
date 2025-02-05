package signup

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"your_project/internal/models"
)

// Service provides methods for user signup.
type Service struct {
	userModel *models.UserModel
}

// NewService creates a new signup service.
func NewService(userModel *models.UserModel) *Service {
	return &Service{
		userModel: userModel,
	}
}

// SignupRequest represents the request payload for user signup.
type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignupResponse represents the response payload for user signup.
type SignupResponse struct {
	Message string `json:"message"`
}

// Signup handles user registration.
func (s *Service) Signup(req SignupRequest) (SignupResponse, error) {
	if req.Email == "" || req.Password == "" {
		return SignupResponse{}, errors.New("email and password are required")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return SignupResponse{}, err
	}

	// Create a new user
	user := &models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// Save the user to the database
	if err := s.userModel.Create(user); err != nil {
		return SignupResponse{}, err
	}

	return SignupResponse{Message: "User registered successfully"}, nil
}
package auth

import (
    "errors"
    "time"

    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
)

// User represents the user model
type User struct {
    ID       string
    Email    string
    Password string
}

// AuthService provides methods for authentication
type AuthService struct {
    secretKey string
}

// NewAuthService creates a new AuthService
func NewAuthService(secretKey string) *AuthService {
    return &AuthService{secretKey: secretKey}
}

// HashPassword hashes the user's password
func (s *AuthService) HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// ValidatePassword checks if the provided password matches the hashed password
func (s *AuthService) ValidatePassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateToken generates a JWT token for the user
func (s *AuthService) GenerateToken(userID string) (string, error) {
    claims := jwt.MapClaims{
        "sub": userID,
        "exp": time.Now().Add(time.Hour * 72).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.secretKey))
}

// ParseToken parses the JWT token and returns the user ID
func (s *AuthService) ParseToken(tokenString string) (string, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(s.secretKey), nil
    })
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims["sub"].(string), nil
    }
    return "", err
}
package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/henriquelazzarino/gookshelf/config"
	"github.com/henriquelazzarino/gookshelf/models"
	"github.com/henriquelazzarino/gookshelf/repositories"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (string, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("invalid credentials1")
	}

	fmt.Println(email, password, user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials2")
	}

	return generateJWT(user.ID, user.Role)
}

func generateJWT(userID string, userRole models.UserRole) (string, error) {
	// Set claims for JWT
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  "gookshelf",                      // Issuer
		"sub":  userID,                           // Subject (user ID)
		"role": userRole,                         // Role
		"exp":  time.Now().Add(time.Hour).Unix(), // Expires in 1 hour
	})

	// Use a secret key for signing
	secretKey := []byte(config.JWTSecret) // Replace with a strong secret key

	return claims.SignedString(secretKey)
}

func VerifyJWT(tokenString string) (*models.User, error) {
	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil // Use the same secret key used for signing
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["sub"].(string)
		user, err := repositories.GetUser(userID)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, errors.New("invalid token")
}

package hashing

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a given password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // Adjust cost as needed
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ComparePassword compares a plain text password with a hashed password
func ComparePassword(plainPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

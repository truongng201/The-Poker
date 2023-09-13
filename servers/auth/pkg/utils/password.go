package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("error hashing password: %v", err)
	}

	return string(hash), nil
}

// CheckPassword checks a password
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

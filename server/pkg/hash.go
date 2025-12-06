// Package pkg provide utilities for application
package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePasswordHash(password string, hashes string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashes), []byte(password))
	return err == nil
}

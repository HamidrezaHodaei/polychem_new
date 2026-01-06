package utils

import (
	"polychem-auth/config"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword - هش کردن پسورد
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password), 
		config.AppConfig.BcryptCost,
	)
	return string(bytes), err
}

// CheckPasswordHash - بررسی صحت پسورد
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
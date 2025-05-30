package utils

import(
	"golang.org/x/crypto/bcrypt"
)

const(
	MinCost int=4 // Minimum cost for bcrypt
	MaxCost int=31 // Maximum cost for bcrypt
	DefaultCost int=10 // Default cost for bcrypt
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
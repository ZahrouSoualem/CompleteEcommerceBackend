package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// this will return the Hashed password of the entered password
func HashPassword(password string) (string, error) {

	Hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost /*the default value is 10*/)

	if err != nil {
		return "", fmt.Errorf("failed to hash a password")
	}

	return string(Hashedpassword), nil
}

func CheckPassword(password string, HashedPassword string) error {

	return bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(password))
}

package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	crypt, err := bcrypt.GenerateFromPassword([]byte(password), 17)
	return string(crypt), err

}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

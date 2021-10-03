package helper

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plain string) (string,error) {
	plaintByte := []byte(plain)
	password,err := bcrypt.GenerateFromPassword(plaintByte,bcrypt.DefaultCost)
	return string(password),err
}

func ComparePassword(plain string, hashed string) error {
	hashedByte := []byte(hashed)
	plainByte := []byte(plain)
	err := bcrypt.CompareHashAndPassword(hashedByte,plainByte)
	if err != nil {
		return errors.New("username / password salah !!!")
	}
	return nil
}
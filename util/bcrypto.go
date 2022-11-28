package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 14
)

func HashPassword(password string) string {
	hashPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		fmt.Println("GenerateFromPassword error: ", err.Error())
		return ""
	}

	return string(hashPasswordBytes)
}

func CheckPasswordHash(hashString, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashString), []byte(password))

	return err == nil
}

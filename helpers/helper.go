package helpers

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func PrintErr(err error, messge string) {
	fmt.Println(messge, err)
}

func PrintMsg(msg string) {
	fmt.Println(msg)
}

func Hash_pass(pass string) (string, error) {

	password := []byte(pass)

	hashedpass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	return string(hashedpass), err

}

func VerifyPassword(hashedPassword, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

}

func GenUuid() string {
	return uuid.New().String()
}
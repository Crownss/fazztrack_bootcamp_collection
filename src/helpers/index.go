package helpers

import (
	"errors"
	"log"
	"math/rand"

	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func Env() error {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("cannot load env file with error:\n", env.Error())
	}
	return env
}

func Admin_Check(models *models.Users) (bool, error) {
	if !models.Is_Admin {
		return false, errors.New("not admin")
	}
	return models.Is_Admin, nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomCode(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func EncryptPassword(s string) (string, error) {
	encrypt, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	return string(encrypt), err
}

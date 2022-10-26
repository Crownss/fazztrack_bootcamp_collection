package helpers

import (
	"errors"
	"log"
	"net/http"

	"github.com/crownss/fazztrack_bootchamp/week_9/src/database/orm/models"
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

var Payload_200 = models.Response{
	Code:    http.StatusOK,
	Message: "Success !",
	Status:  http.StatusText(http.StatusOK),
}
var Payload_500 = models.Response{
	Code:    http.StatusInternalServerError,
	Message: "Error !",
	Status:  http.StatusText(http.StatusInternalServerError),
}
var Payload_400 = models.Response{
	Code:    http.StatusBadRequest,
	Message: "Error !",
	Status:  http.StatusText(http.StatusBadRequest),
}

func EncryptPassword(s string) (string, error) {
	encrypt, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	return string(encrypt), err
}

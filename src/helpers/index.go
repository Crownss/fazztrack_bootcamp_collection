package helpers

import (
	"log"
	"net/http"

	"github.com/crownss/fazztrack_bootchamp/week_8/src/database/orm/models"
	"github.com/joho/godotenv"
)

func Env(file string) error {
	env := godotenv.Load(file)
	if env != nil {
		log.Fatal("cannot load env file with error:\n", env.Error())
	}
	return env
}


var Payload_200 = models.Response{
	Code: http.StatusOK,
	Message: "Success !",
	Status: http.StatusText(http.StatusOK),
}
var Payload_500 = models.Response{
	Code: http.StatusInternalServerError,
	Message: "Error !",
	Status: http.StatusText(http.StatusInternalServerError),
}
var Payload_400 = models.Response{
	Code: http.StatusBadRequest,
	Message: "Error !",
	Status: http.StatusText(http.StatusBadRequest),
}
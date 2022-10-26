package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crownss/fazztrack_bootchamp/week_8/src/helpers"
	"github.com/crownss/fazztrack_bootchamp/week_8/src/routers"
)

func main() {
	helpers.Env(".env")
	mainRoute, err := routers.New()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("service run on port 8000")
	http.ListenAndServe(":8000", mainRoute)
}

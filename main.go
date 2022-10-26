package main

import (
	"log"
	"os"

	"github.com/crownss/fazztrack_bootchamp/week_10/src/config"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := config.Run(os.Args[1:]); err != nil {
		log.Fatalln(err.Error())
	}
}

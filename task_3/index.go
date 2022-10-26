package task3

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/crownss/fazztrack_bootchamp/week_7/helpers"
)

func Ticket(code string, c chan<- string) {
	if len(code) < 5 {
		fmt.Println("code movie must be greater than 5")
	}
	fmt.Println("Movie ticket:", code)
	c <- code
	helpers.WG.Done()
}

func PopCorn(c <-chan string) {
	rand.Seed(time.Now().UnixNano())
	value := <-c
	random := rand.Intn(len(value))
	if random == 0 {
		random = 1
	}
	fmt.Println("you get", random, "free small popcorn for code movie", value)
	helpers.WG.Done()
}
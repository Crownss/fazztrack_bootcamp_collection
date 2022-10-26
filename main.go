package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("reversal triangle")
	n := Methode{7}
	n.ReversalTriangle()
	fmt.Println("\n\ngenerate password")
	rand.Seed(time.Now().UnixNano())
	GenPassword("jesen123", "strong")
	fmt.Println("\n\nfilm and flight")
	fmt.Println(n.Flight())
}

type Methode struct {
	N int
}

func (methode *Methode) ReversalTriangle() {
	for i := methode.N; i >= 1; i-- {
		for space := 1; space <= methode.N-i; space++ {
			fmt.Print("  ")
		}
		for j := i; j <= 2*i-1; j++ {
			fmt.Printf("* ")
		}
		for j := 0; j < i-1; j++ {
			fmt.Printf("* ")
		}
		fmt.Println("")
	}
}

func GenPassword(password, level string) string {
	randomStrings := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ2134567890./?,<>;:'{}[]()-=+~`@"
	letter := strings.Split(password, "")
	b := make([]byte, len(password))
	if strings.ToLower(level) == "low" {
		if len(password) < 6 {
			for i := range letter {
				b[i] = randomStrings[rand.Intn(len(randomStrings))]
				letter[i] = string(b)
			}
			fmt.Println(strings.Join(letter, ""))
			return strings.Join(letter, "")
		}
		for i := range letter {
			b[i] = password[rand.Intn(len(password))]
		}
		fmt.Println(string(b))
		return string(b)
	} else if strings.ToLower(level) == "medium" || strings.ToLower(level) == "med" {
		if len(password) < 6 {
			for i := range letter {
				b[i] = password[rand.Intn(len(letter))]
				letter[i] = string(b)
			}
			fmt.Println(strings.Join(letter, ""))
			return strings.Join(letter, "")
		}
		for i := range letter {
			b[i] = password[rand.Intn(len(letter))]
		}
		fmt.Println(string(b))
		return string(b)
	} else if strings.ToLower(level) == "strong" {
		if len(password) < 6 {
			for i := range letter {
				b[i] = randomStrings[rand.Intn(len(randomStrings))]
				letter[i] = string(b)
			}
			fmt.Println(strings.Join(letter, ""))
			return strings.Join(letter, "")
		}
		for i := range letter {
			b[i] = randomStrings[rand.Intn(len(randomStrings))]
		}
		fmt.Println(string(b))
		return string(b)
	}
	return string(b)
}

func (methode *Methode) Flight() string {
	var result string
	datas := []int{1, 7, 3, 4, 8, 9}
	for i, first := range datas {
		for j, second := range datas {
			if j != i && first+second == methode.N {
				result = fmt.Sprintf("first film %d and %d", first, second)
				return result
			} else if j != i && first+second != methode.N {
				result = fmt.Sprintf("Error duration flight %d is not matched with any duration film", methode.N)
			}
		}
	}
	return result
}

package task1

import (
	"fmt"

	"github.com/crownss/fazztrack_bootchamp/week_7/helpers"
)

type DeretBilangan struct {
	N int
}

func (db *DeretBilangan) Prima() []int {
	if db.N <= 2 {
		fmt.Println("number must greater than 2")
	}
	var result []int
	for i := 1; i <= db.N; i++ {
		if i%2 == 1 || i == 2 {
			result = append(result, i)
		}
	}
	// fmt.Println("bilangan prima:\n")
	helpers.WG.Done()
	return result
}

func (db *DeretBilangan) Ganjil() []int {
	if db.N <= 2 {
		fmt.Println("number must greater than 2")
	}
	var result []int
	for i := 0; i <= db.N; i++ {
		if i%2 == 1 {
			result = append(result, i)
		}
	}
	// fmt.Println("bilangan ganjil:\n")
	helpers.WG.Done()
	return result
}

func (db *DeretBilangan) Genap() []int {
	if db.N <= 2 {
		fmt.Println("number must greater than 2")
	}
	var result []int

	for i := 0; i <= db.N; i++ {
		if i%2 == 0 {
			result = append(result, i)
		}
	}
	// fmt.Println("bilangan genap:\n")
	helpers.WG.Done()
	return result
}

func (db *DeretBilangan) Fibonanci() []int {
	if db.N <= 2 {
		fmt.Println("number must greater than 2")
	}
	first, second := 0, 1
	var third int
	var result []int

	for i := 1; i <= db.N; i++ {
		if third > db.N {
			break
		}
		third = first
		first = second
		second = third + first
		result = append(result, third)
	}
	helpers.WG.Done()
	return result
}

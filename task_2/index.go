package task2

import (
	"fmt"
)

type DeretBilangan struct {
	N int
}

func (db *DeretBilangan) Prima() int {
	if db.N <= 2 {
		fmt.Println("number must greater than 2")
	}
	for i := 1; i <= db.N; i++ {
		if i%2 == 1 || i == 2 {
			fmt.Println("bilangan prima\n", i)
		}
	}
	return db.N
}

func (db *DeretBilangan) Ganjil() int {
	if db.N <= 2 {
		fmt.Println("number must greater than 2")
	}
	for i := 0; i <= db.N; i++ {
		if i%2 == 1 {
			fmt.Println("bilangan ganjil\n", i)
		}
	}
	return db.N
}

func (db *DeretBilangan) Genap() int {
	if db.N <= 2 {
		fmt.Println("number must greater than 2")
	}
	for i := 0; i <= db.N; i++ {
		if i%2 == 0 {
			fmt.Println("bilangan genap\n", i)
		}
	}
	return db.N
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
	return result
}

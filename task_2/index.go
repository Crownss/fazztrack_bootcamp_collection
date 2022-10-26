package task2

import (
	"fmt"

	"github.com/crownss/fazztrack_bootchamp/week_7/helpers"
)

// sender
func Fibo(value int, ch chan<- []int) {
	first, second := 0, 1
	var third int
	var result []int

	for i := 1; i <= value; i++ {
		if third > value {
			break
		}
		third = first
		first = second
		second = third + first
		result = append(result, third)
	}
	fmt.Println("Fibonanci (task 2):", result)
	ch <- result
	helpers.WG.Done()
}

// receiver
func OddEven(ch <-chan []int) {
	value := <-ch
	var ganjil []int
	var genap []int
	for i, v := range value {
		if value[i]%2 == 0 {
			genap = append(genap, v)
		}
		if value[i]%2 == 1 {
			ganjil = append(ganjil, v)
		}
	}
	fmt.Println("Ganjil (task 2):", ganjil)
	fmt.Println("Genap (task 2):", genap)
	helpers.WG.Done()
}

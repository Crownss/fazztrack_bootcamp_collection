package main

import (
	"fmt"

	task3 "github.com/crownss/fazztrack_bootchamp/week_6/task_3"
)

func main() {
	// fmt.Println("pembulatan")
	// task1.Pembulatan(40.99)

	// fmt.Println("deret bilangan")
	// v := task2.DeretBilangan{N: 7}
	// fmt.Println(v.Prima())
	// fmt.Println()
	// fmt.Println(v.Ganjil())
	// fmt.Println()
	// fmt.Println(v.Genap())
	// fmt.Println()
	// fmt.Println(v.Fibonanci())
	// fmt.Println()
	// fmt.Println()

	// var inter task3.Hitung
	// inter = &task3.Persegi{7}
	// fmt.Printf("value %v", inter)

	v2 := task3.Persegi{S: 7}
	fmt.Println(v2.Luas())
	fmt.Println(v2.Keliling())
	fmt.Println(v2.Volume())
}

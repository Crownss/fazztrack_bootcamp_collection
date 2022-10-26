package main

import (
	"github.com/crownss/fazztrack_bootchamp/week_7/helpers"
	// task1 "github.com/crownss/fazztrack_bootchamp/week_7/task_1"
	// task2 "github.com/crownss/fazztrack_bootchamp/week_7/task_2"
	task3 "github.com/crownss/fazztrack_bootchamp/week_7/task_3"
	// "github.com/crownss/fazztrack_bootchamp/week_7/tryit"
)

func main() {
	helpers.WG.Add(2)
	//Try it
	// a := []int{7, 10, 2, 34, 33, -12, -8, 4}
	// ch := make(chan int)
	// go tryit.Sum(a[:len(a)/2], ch)
	// go tryit.Sum(a[len(a)/2:], ch)
	// res := <-ch
	// fmt.Println(res)

	//Task 1
	// b := task1.DeretBilangan{N: 10}
	// go func() { fmt.Println("Prima (task 1):", b.Prima()) }()
	// go func() { fmt.Println("Ganjil (task 1):", b.Ganjil()) }()
	// go func() { fmt.Println("Genap (task 1):", b.Genap()) }()
	// go func() { fmt.Println("Fibonanci (task 1):", b.Fibonanci()) }()

	//Task 2
	// ch1 := make(chan []int)
	// go task2.Fibo(7, ch1)
	// go task2.OddEven(ch1)

	// //Task 3
	ch2 := make(chan string, 5)
	go task3.Ticket("drss1f3", ch2)
	go task3.PopCorn(ch2)
	helpers.WG.Wait()
	// close(ch1)
	close(ch2)
}

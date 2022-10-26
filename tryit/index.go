package tryit

func Sum(d []int, ch chan<- int) {
	var result int
	for i := range d {
		result += d[i] + d[i]
	}
	ch <- result
}

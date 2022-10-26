package task3

import "math"

type Hitung2d interface {
	Luas() float64
	Keliling() float64
}

type Hitung3d interface {
	Volume() float64
}

type Hitung interface {
	Hitung2d
	Hitung3d
}

type Persegi struct {
	S int
}

func (sisi *Persegi) Luas() float64 {
	return math.Pow(float64(sisi.S), 2)
}

func (sisi *Persegi) Keliling() float64 {
	return float64(sisi.S) * 4
}

func (sisi *Persegi) Volume() float64 {
	return math.Pow(float64(sisi.S), 3)
}

package golang_united_school_homework

import "math"

type Circle struct {
	Radius float64
}

func (b Circle) CalcArea() float64 {
	s := math.Pi * math.Pow(b.Radius, 2)
	return s
}

func (b Circle) CalcPerimepter() float64 {
	p := 2 * b.Radius * math.Pi
	return p
}




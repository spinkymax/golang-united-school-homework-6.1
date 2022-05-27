package golang_united_school_homework

import "math"

type Triangle struct {
	Side float64
}

func (b Triangle) CalcArea() float64 {
	s := b.Side * b.Side * math.Sqrt(3) / 4
	return s
}

func (b Triangle) CalcPerimepter() float64 {
	p := 3 * b.Side
	return p
}

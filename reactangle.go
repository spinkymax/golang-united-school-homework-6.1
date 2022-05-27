package golang_united_school_homework

type Rectangle struct {
	Height, Weight float64
}

func (b Rectangle) CalcArea() float64 {
	s := b.Height * b.Weight
	return s
}

func (b Rectangle) CalcPerimepter() float64 {
	p := 2 * (b.Height + b.Weight)
	return p
}



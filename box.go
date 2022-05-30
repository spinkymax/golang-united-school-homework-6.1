package golang_united_school_homework

import (
	"errors"
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) > b.shapesCapacity {
		return errors.New("Sorry, but capacity goes out of range")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) < i || i < 0 {
		return nil, errors.New("Sorry, but shape by index doesn't exist or index goes out of the range")
	}
	return b.shapes[i], nil

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) < i || i < 0 {
		return nil, errors.New("Sorry, but shape by index doesn't exist or index goes out of the range")
	}

	b.shapes[i] = b.GetByIndex[len(b.GetByIndex)-1]
	return b.shapes[i], nil

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if len(b.shapes) < i || i < 0 {
		return nil, errors.New("Sorry, but shape by index doesn't exist or index goes out of the range")
	}
	recoverShape := b.GetByIndex
	return recoverShape, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64 = 0
	for _, number := range b.shapes {
		sum += number.CalcPerimeter
	}
	return sum

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64 = 0
	for _, number := range b.shapes {
		sum += number.CalcArea
	}
	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	round := 0
	if round == 0 {
		return fmt.Errorf("Sorry, but circles are not exist in the list")
	}
	for i := len(b.shapes) - 1; i >= 0; i-- {
		if _, ok := b.shapes[i].(*Circle); ok {
			round++
			b.shapes[i] = b.GetByIndex[len(b.GetByIndex)-b.shapes[i].(*Circle)]
		}
	}
	return nil


}

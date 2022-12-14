package calculator

import (
	"github.com/thienkb1123/solid-principles-in-golang/src/srp/after/shapes"
)

type Calculator struct {
	square *shapes.Square
}

func NewSquare(square *shapes.Square) *Calculator {
	return &Calculator{square}
}

func (c *Calculator) AreaOfSquare() float64 {
	return c.square.Width * c.square.Height
}

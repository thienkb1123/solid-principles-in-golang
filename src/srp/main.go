package srp

import (
	"github.com/thienkb1123/solid-principles-in-golang/src/srp/after/calculator"
	"github.com/thienkb1123/solid-principles-in-golang/src/srp/after/shapes"
)

func Run() {
	square := shapes.NewSquare(2, 3)
	cal := calculator.NewSquare(square)
	result := cal.AreaOfSquare()

	calculator.PrintSquareArea(result)
}

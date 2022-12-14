package befor

import "fmt"

type Square struct {
	Width  float64
	Height float64
}

func (s *Square) MakeSquare(width, heigth float64) {
	s.Width = width
	s.Height = heigth
}

func (s *Square) GetWidth() float64 {
	return s.Width
}

func (s *Square) GetHeight() float64 {
	return s.Height
}

func (s *Square) AreaOfSquare() float64 {
	return s.Width * s.Height
}

func (s *Square) PrintSquareArea(result float64) {
	fmt.Printf("square area: %f\n", result)
}

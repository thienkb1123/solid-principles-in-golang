package shapes

type Square struct {
	Width  float64
	Height float64
}

func NewSquare(width, heigth float64) *Square {
	return &Square{
		Width:  width,
		Height: heigth,
	}
}

func (s *Square) GetWidth() float64 {
	return s.Width
}

func (s *Square) GetHeight() float64 {
	return s.Height
}

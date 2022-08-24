package figuras

import "math"

type Circulo struct {
	Radio int
}

func (cl *Circulo) area() float64 {
	pi := math.Pi
	area := pi * (float64(cl.Radio) * float64(cl.Radio))
	return area
}

func (cl *Circulo) perimetro() float64 {
	pi := math.Pi
	perimetro := 2 * pi * float64(cl.Radio)
	return perimetro
}

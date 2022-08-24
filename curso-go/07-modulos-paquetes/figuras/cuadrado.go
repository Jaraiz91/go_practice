package figuras

type Cuadrado struct {
	Ancho  float64
	Altura float64
}

func (cd *Cuadrado) area() float64 {
	area := cd.Ancho * cd.Altura
	return area
}

func (cd *Cuadrado) perimetro() float64 {
	perimetro := 2*cd.Ancho + (2 * cd.Altura)
	return float64(perimetro)
}

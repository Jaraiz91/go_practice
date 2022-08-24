package figuras

import "fmt"

type Forma interface {
	area() float64
	perimetro() float64
}

func Medidas(forma Forma) {
	fmt.Println("Medidas:", forma)
	fmt.Println("Área:", forma.area())
	fmt.Println("Perímetro:", forma.perimetro())
}

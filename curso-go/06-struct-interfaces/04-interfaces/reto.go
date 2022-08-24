package main

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
	perimetro() float64
}

type Cuadrado struct {
	ancho  float64
	altura float64
}

type Circulo struct {
	radio int
}

func (cd *Cuadrado) area() float64 {
	area := cd.ancho * cd.altura
	return area
}

func (cl *Circulo) area() float64 {
	pi := math.Pi
	area := pi * (float64(cl.radio) * float64(cl.radio))
	return area
}

func (cd *Cuadrado) perimetro() float64 {
	perimetro := 2*cd.ancho + (2 * cd.altura)
	return float64(perimetro)
}

func (cl *Circulo) perimetro() float64 {
	pi := math.Pi
	perimetro := 2 * pi * float64(cl.radio)
	return perimetro
}

func calcularArea(forma Forma) {
	fmt.Println("Area:", forma.area())
}

func calcularPerimetro(forma Forma) {
	fmt.Println("Perimetro:", forma.perimetro())
}

func main() {
	cuad := Cuadrado{ancho: 2, altura: 4}
	circ := Circulo{radio: 3}
	calcularArea(&cuad)
	calcularPerimetro(&cuad)
	calcularArea(&circ)
	calcularPerimetro(&circ)

}

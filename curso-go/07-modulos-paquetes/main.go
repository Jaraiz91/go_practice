package main

import (
	"fmt"
	"pkg/figuras"
	"pkg/models"

	"github.com/donvito/hellomod"
)

func main() {
	cua1 := figuras.Cuadrado{Ancho: 4, Altura: 2}
	circ1 := figuras.Circulo{Radio: 3}
	figuras.Medidas(&cua1)
	figuras.Medidas(&circ1)

	p1 := models.Persona{}
	p1.Constructor("Javier", 31)
	fmt.Println(p1.GetNombre())
	p1.SetNombre("Isabel")
	fmt.Println(p1.GetNombre())
	fmt.Println("Probando paquete de terceros:")
	hellomod.SayHello()

}

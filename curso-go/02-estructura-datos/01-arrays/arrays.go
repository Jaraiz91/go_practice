package main

import "fmt"

func main() {
	//arrays
	var numeros [5]int
	numeros[0] = 10
	numeros[1] = 20
	fmt.Println(numeros)

	//arrays con valores
	nombres := [3]string{"Javier", "Pablo", "Isabel"}
	fmt.Println(nombres[2])

	//arrays sin límite
	colores := [...]string{
		"rojo",
		"azul",
		"amarillo",
		"verde",
	}
	fmt.Println(colores, "num colores: ", len(colores))

	//indices definidos
	monedas := [...]string{
		0: "dolar",
		2: "euro",
		3: "libra",
	}
	//len va a ser 4, porque nos hemos saltado la posicion 1
	fmt.Println(monedas, len(monedas))

	submonedas := monedas[:2]
	fmt.Println(submonedas)
}

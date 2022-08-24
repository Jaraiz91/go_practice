package main

import "fmt"

func sumar(nombre string, numeros ...int) (string, int) {
	mensaje := fmt.Sprintf("la suma de %s es:", nombre)
	var total int
	for _, num := range numeros {

		total += num
	}
	return mensaje, total
}

func main() {
	mensaje, result := sumar("Javier", 4, 5, 6, 8)
	fmt.Println(mensaje, result)
}

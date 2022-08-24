package main

import "fmt"

func modificarValor(cadena *string) {
	fmt.Printf("%p\n", cadena)
	*cadena = "Hola desde la función"
}

func main() {
	cadena := "Hola mundo de Go"
	fmt.Printf("%p\n", &cadena)
	fmt.Println("Antes de la función:", cadena)

	modificarValor(&cadena) //referencia

	fmt.Println("Después de la función:", cadena)
}

package main

import "fmt"

func main() {
	hola := "hola"
	mundo := "mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Javier"
	edad := 31

	fmt.Printf("Hola, %s y su edad es %d \n", nombre, edad)

	mensaje := fmt.Sprintf("Hola, %v y su edad es %v", nombre, edad)
	fmt.Println(mensaje)

	fmt.Printf("Nombre: %T \n", nombre)
	fmt.Print("Ingrese otro nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("otro nombre", nombre)

}

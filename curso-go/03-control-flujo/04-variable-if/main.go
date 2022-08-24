package main

import "fmt"

func main() {
	if nombre, edad := "Isabel", 31; nombre == "Javier" {
		fmt.Printf("Hola %v \n", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad %d \n", nombre, edad)
	}

	//Obtener valor de mapa
	users := make(map[string]string)
	users["Javier"] = "jaraiz@gmail.com"
	users["Isabel"] = "iromerojordano@gmail.com"

	if correo, ok := users["Javier"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}

}

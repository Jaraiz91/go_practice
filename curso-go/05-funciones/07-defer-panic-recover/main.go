package main

import (
	"fmt"
	"os"
)

func main() {
	//Función
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Oops al parecer el programa no finalizó de forma correcta")
		}
	}()

	if file, err := os.Open("hoa.txt"); err != nil {
		panic("Error al abrir el archivo")
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo!")
			file.Close()
		}()
		contenido := make([]byte, 254)
		longitud, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:longitud])
		fmt.Println(contenidoArchivo)
	}

}

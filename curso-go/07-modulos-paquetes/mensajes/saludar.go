package mensaje

import "fmt"

func Hola() {
	fmt.Println("Hola desde el paquete mensaje")
}

const mensaje = "Hola desde mi constante"

func Imprimir() {
	fmt.Println(mensaje)
	funcprivada()
}

func funcprivada() {
	fmt.Println("Hola desde la función privada")
}

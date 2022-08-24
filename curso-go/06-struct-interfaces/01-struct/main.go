package main

import "fmt"

// struct persona
type Persona struct {
	nombre string
	edad   int
}

// metodos

func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\nEdad: %v\n", p.nombre, p.edad)
}

func (p *Persona) editarNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

// Herencia
type Empleado struct {
	Persona
	sueldo float64
}

//crear objetos desde esa estructura

func main() {
	p1 := Persona{"Javier", 31}
	fmt.Println(p1)
	p1.nombre = "Isabel"
	fmt.Println(p1)
	p2 := Persona{nombre: "Pablo", edad: 26}
	fmt.Println(p2)
	p1.imprimir()
	p2.imprimir()
	p1.editarNombre("Javier")
	p1.imprimir()
	empleado1 := Empleado{sueldo: 500}
	empleado1.nombre = "Pedro"
	empleado1.edad = 32
	empleado1.imprimir()
	fmt.Println(empleado1)

}

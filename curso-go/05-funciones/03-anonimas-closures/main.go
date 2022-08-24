package main

import (
	"fmt"
	"strings"
)

// closures
func repeat(n int) func(cadena string) string {
	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}
func main() {
	//funcion anonima
	func() {
		fmt.Println("Hola desde la función anónima")
	}()
	repeat3 := repeat(3)
	fmt.Println(repeat3("Hola"))
}

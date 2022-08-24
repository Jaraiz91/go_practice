package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) int {
	//4+4+5+8
	valores := strings.Split(expresion, "+")
	suma := 0
	for i := 0; i <= len(valores)-1; i++ {
		num, err := strconv.Atoi(valores[i])
		fmt.Println(err)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error: tiene que ingresar un numero entero y solamente con el operador de suma")
		} else {
			suma = suma + num
		}
	}
	return suma
}

func main() {
	var expresion string
	var result int
	fmt.Print("=>")
	fmt.Scanln(&expresion)
	result = sumar(expresion)
	fmt.Println(result)

}

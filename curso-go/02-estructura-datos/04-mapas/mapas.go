package main

import (
	"fmt"
)

func main() {
	dias := make(map[int]string)

	fmt.Println(dias)

	//agregar datos
	dias[1] = "Domingo"
	dias[2] = "lunes"
	dias[3] = "martes"
	dias[4] = "miercoles"
	dias[5] = "jueves"
	dias[6] = "viernes"
	dias[7] = "sabado"

	fmt.Println(dias)

	dias[7] = "SABADO"

	fmt.Println(dias)

	delete(dias, 1)

	fmt.Println(dias)

	estudiantes := make(map[string][]int)
	estudiantes["javier"] = []int{12, 21, 4}
	estudiantes["Isabel"] = []int{13, 22, 5}

	fmt.Println(estudiantes)

}

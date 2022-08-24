package main

import "fmt"

func saludar(nombre string) {
	fmt.Printf("Hello %s \n", nombre)
}

func sumar(a, b int) int {
	return a + b
}

func get_user_numbers() int {
	var a int
	fmt.Print("ingrese numero: ")
	fmt.Scanln(&a)
	return a
}

func divide(a, b int) int {
	return a / b
}

func module(a, b int) int {
	return a % b
}

func exercise_01() {
	// func to execute exercise one. user is asked for two numbers and the added result will be returned
	a := get_user_numbers()
	b := get_user_numbers()
	r := sumar(a, b)
	fmt.Println("suma: ", r)
}

func exercise_02() {
	a := get_user_numbers()
	b := get_user_numbers()
	cociente := divide(a, b)
	resto := module(a, b)
	fmt.Println("cociente: ", cociente)
	fmt.Println("resto: ", resto)
}

func exercise_03() {
	precio := get_user_numbers()
	var precio_final float64 = float64(precio) * 1.21
	fmt.Println("precio final: ", precio_final)

}

func main() {
	saludar("Javier")
	exercise_01()
	exercise_02()
	exercise_03()
}

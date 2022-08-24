package main

import "fmt"

func main() {
	//slicing, son mutables a diferencia de los array
	numeros := []int{1, 2, 3}
	numeros = append(numeros, 4)
	fmt.Println(numeros)

	//sub slicing
	subslicing := numeros[:2]
	subslicing[0] = 100

	fmt.Println(subslicing)
	fmt.Println(numeros)

	meses := []string{"enero", "febrero", "marzo"}

	fmt.Printf("len: %v, cap: %v, %p", len(meses), cap(meses), meses)
}

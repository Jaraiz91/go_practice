package main

import "fmt"

func main() {
	a := 4
	b := 5
	var r bool

	//igualdad
	r = a == b
	fmt.Println(r)
	fmt.Printf("%v es igual que %v? %v", a, b, r)
}

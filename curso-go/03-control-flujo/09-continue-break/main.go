package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 5 {
			fmt.Println("salto el 5")
			continue
		}

		if i == 8 {
			fmt.Println("rompo bucle")
			break
		}
		fmt.Println(i)
	}
}

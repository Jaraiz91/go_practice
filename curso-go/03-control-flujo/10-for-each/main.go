package main

import "fmt"

func main() {
	nombres := [...]string{"javier", "Isabel", "Pablo"}

	/*for i := 0; i < len(nombres); i++ {
		fmt.Println(nombres[i])
	}*/
	for _, elemento := range nombres {
		fmt.Println(elemento)
	}

	for indice, elemento := range nombres {
		fmt.Println(indice, elemento)
	}
}

package main

import (
	"errors"
	"fmt"
)

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No es posible dividir entre 0")
	} else {
		return dividendo / divisor, nil
	}

}

func main() {
	result, err := division(4, 0)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

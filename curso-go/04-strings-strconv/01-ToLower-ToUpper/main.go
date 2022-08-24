package main

import (
	"fmt"
	"strings"
)

func reverse(cadena string) string {
	arrayCadena := strings.Split(cadena, "")
	arrayInvertida := make([]string, 0)

	for i := len(arrayCadena) - 1; i >= 0; i-- {
		arrayInvertida = append(arrayInvertida, arrayCadena[i])
	}

	return strings.Join(arrayInvertida, "")
}

func es_palindromo(palabra string) bool {
	palabra = strings.ToLower(palabra)
	palabra = strings.ReplaceAll(palabra, " ", "")
	palabra_invertida := reverse(palabra)
	fmt.Println("palabra invertida: ", palabra_invertida)
	return palabra == palabra_invertida

}
func main() {
	if es_palindromo("luz azul") {
		fmt.Println("Es palíndromo")
	} else {
		fmt.Println("No es palíndromo")
	}

}

package main

import "fmt"

func main() {
	/** App para restaurante
	* Descuentos de 10% hasta 100 de consumo
	* 20% para 200 de consumo
	* 30% para 300 de consumo
	iva 21%
	*/
	var consumo, descuento float64
	var datosDescuento string

	//Entrada de datos
	fmt.Println("Ingrese consumo: ")
	fmt.Scanln(&consumo)

	if consumo >= 0 {
		if consumo <= 100 {
			descuento = consumo * 0.1
			datosDescuento = "10%"
		} else if consumo > 100 && consumo <= 200 {
			descuento = consumo * 0.2
			datosDescuento = "20%"
		} else if consumo > 200 {
			datosDescuento = "30%"
			descuento = consumo * 0.3
		} else {
			fmt.Println("error al ingresar el importe")
		}
	}
	iva := (consumo - descuento) * 0.21
	fmt.Println("-------FACTURA DE CONSUMO-------")
	fmt.Println("total a pagar: ", (consumo-descuento)+iva)
	fmt.Println("consumo", consumo)
	fmt.Println("descuento: ", datosDescuento)
	fmt.Println("IVA", iva)
}

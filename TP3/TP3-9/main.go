package main

import (
	"fmt"
)

func main() {

	const iva = 0.21

	var (
		cant    int
		pre     float64
		preNeto float64
		total   float64
		iVA     float64
		iVAT    float64
	)

	fmt.Print("Ingrese la cantidad de objetos:")
	fmt.Scanln(&cant)
	fmt.Print("Ingrese el precio unitario del objeto:")
	fmt.Scanln(&pre)

	//Precio neto
	preNeto = float64(cant) * pre

	//Precio con iva
	iVA = pre * iva
	iVAT = pre + iVA

	//Precio precioTotal
	total = preNeto + (iVA * float64(cant))

	fmt.Printf("el precio neto es de %v, el precio del producto con IVA es de %v y el importe total es de %v", preNeto, iVAT, total)
}

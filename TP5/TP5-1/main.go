package main

import "fmt"

func main() {
	var (
		pes  float64
		dol  float64
		coti float64
	)

	fmt.Print("ingrese la cotizacion del dia:")
	fmt.Scanln(&coti)
	fmt.Print("Ingrese la cantidad de pesos con los que desea comprar los dolares:")
	fmt.Scanln(&pes)
	for pes != 0 {
		fmt.Print("ingrese la cantidad de dolares que desea comprar:")
		fmt.Scanln(&dol)
		opValid(pes, dol, coti)
		fmt.Print("Ingrese la cantidad de pesos con los que desea comprar los dolares:")
		fmt.Scanln(&pes)
	}
}

func opValid(pes, dol, coti float64) (resultado bool) {
	if (pes / coti) >= dol {
		resultado = true
		fmt.Println("Puedes comprar dolares")
	} else {
		resultado = false
		fmt.Println("No puedes comprar dolares")
	}

	return resultado
}

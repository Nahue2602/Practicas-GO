package main

import (
	"fmt"
)

func main() {
	var excursiones int
	cantA := 0
	cantB := 0
	cantC := 0

	for i := 1; i <= 10; i++ {
		fmt.Printf("Ingrese la cantidad de excursiones vendidas por el promotor %d: ", i)
		fmt.Scanln(&excursiones)

		categoria := categoriaPromotor(excursiones)
		importe := importeAPagar(excursiones)

		switch categoria {
		case "A":
			cantA++
		case "B":
			cantB++
		case "C":
			cantC++
		}

		fmt.Printf("Promotor %d - Categoría: %s - Importe a pagar: $%d\n", i, categoria, importe)
	}

	fmt.Println("\nResumen:")
	fmt.Printf("Cantidad de promotores en Categoría A: %d\n", cantA)
	fmt.Printf("Cantidad de promotores en Categoría B: %d\n", cantB)
	fmt.Printf("Cantidad de promotores en Categoría C: %d\n", cantC)
}

// Función que devuelve la categoría de un promotor
func categoriaPromotor(cant int) string {
	if cant < 10 {
		return "A"
	} else if cant >= 10 && cant < 100 {
		return "B"
	} else {
		return "C"
	}
}

// Función que devuelve el importe total a pagar, usando la categoría
func importeAPagar(cant int) int {
	categoria := categoriaPromotor(cant)
	switch categoria {
	case "A":
		return cant * 1000
	case "B":
		return cant * 1500
	case "C":
		return cant * 1900
	default:
		return 0
	}
}

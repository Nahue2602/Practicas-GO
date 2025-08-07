package main

import "fmt"

//Funcion principal
func main() {
	var (
		ventas     int
		count      int = 10
		cA, cB, cC int
	)
	for i := 0; i < count; i++ {
		fmt.Print("Ingrese la cantidad de ventas realizadas:")
		fmt.Scanln(&ventas)
		categoria := categoriaPromotor(ventas)
		importe := importeAPagar(ventas)

		switch categoria {
		case "A":
			cA++
		case "B":
			cB++
		case "C":
			cC++
		}
		fmt.Println("El vendedor es clase", categoria, "y cobrara", importe)
	}

	fmt.Println("la cantidad de vendedores de clase A es de:", cA)
	fmt.Println("la cantidad de vendedores de clase B es de:", cB)
	fmt.Println("la cantidad de vendedores de clase C es de:", cC)

}

func categoriaPromotor(cant int) string {
	if cant < 10 {
		return "A"
	} else if cant >= 10 && cant < 100 {
		return "B"
	} else {
		return "C"
	}
}

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

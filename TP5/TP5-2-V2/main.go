package main

import "fmt"

func main() {
	var (
		numero int
		count  int = 5
		verif1 bool
		verif2 bool
	)

	for i := 0; i < count; i++ {
		fmt.Println("Ingrese un Numero: ")
		fmt.Scanln(&numero)

		verif1 = verificarMultiplo(numero, 3)
		verif2 = verificarMultiplo(numero, 5)

		if verif1 && !verif2 {
			fmt.Println("Cumple con las Condiciones")
		} else {
			fmt.Println("No cumple con las Condiciones")
		}
	}
}

func verificarMultiplo(numero, multi int) bool {
	varAuxiliar := numero

	for varAuxiliar > 0 {
		varAuxiliar = varAuxiliar - multi
	}

	return varAuxiliar == 0
}

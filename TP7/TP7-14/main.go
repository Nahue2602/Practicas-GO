package main

import "fmt"

const (
	num1 = 3
	num2 = 3
	num3 = 4
)

func main() {
	var (
		A [num1][num2]int
		B [num3]int
	)

	for f := 0; f < num1; f++ {
		for c := 0; c < num2; c++ {
			fmt.Printf("Ingrese el numero de la fila %v y la columna %v:", f+1, c+1)
			fmt.Scanln(&A[f][c])
		}
	}

	for p := 0; p < num3; p++ {
		fmt.Printf("Ingrese el numero de la posicion %v:", p+1)
		fmt.Scanln(&B[p])
	}

	total := verificarSuma(A, B)
	fmt.Print("La suma de los elementos de la matriz A que no se repiten en el vector B es de :", total)
}

func verificarSuma(A [num1][num2]int, B [num3]int) (suma int) {
	for f := 0; f < num1; f++ {
		for c := 0; c < num2; c++ {
			auxiliar := 0
			for p := 0; p < num3; p++ {
				if A[f][c] == B[p] {
					auxiliar = 0
					break
				} else {
					auxiliar = A[f][c]
				}
			}
			suma += auxiliar
		}
	}
	return suma
}

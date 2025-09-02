package main

import "fmt"

const (
	num1 int = 5
	num2 int = 3
)

func main() {
	var (
		a [num1][num2]int
	)

	for f := 0; f < num1; f++ {
		for c := 0; c < num2; c++ {
			fmt.Print("Ingrese un numero:")
			fmt.Scanln(&a[f][c])
		}
	}

	menor, fila, columna := verificarMenor(a)
	fmt.Printf("El numero mas chico es %v, este esta ubicado en la fila %v y en la columna %v", menor, fila, columna)
}

func verificarMenor(a [num1][num2]int) (menor, fila, columna int) {
	for f := 0; f < num1; f++ {
		for c := 0; c < num2; c++ {
			if a[f][c] == a[0][0] {
				menor = a[f][c]
				fila = f + 1
				columna = c + 1
			} else if menor > a[f][c] {
				menor = a[f][c]
				fila = f + 1
				columna = c + 1
			}
		}
	}
	return menor, fila, columna
}

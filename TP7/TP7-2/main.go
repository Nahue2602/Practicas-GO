package main

import "fmt"

const (
	num1 int = 3
	num2 int = 2
)

func main() {
	var (
		a [num1][num2]int
	)

	for f := 0; f < num1; f++ {
		for c := 0; c < num2; c++ {
			fmt.Print("Ingrese 30 numeros:")
			fmt.Scanln(&a[f][c])
		}
	}

	sumaDeElementos, cantElementos := suma(a)
	fmt.Printf("La suma de los elementos es %v y la cantidad de elementos utilizados es de %v", sumaDeElementos, cantElementos)
}

func suma(a [num1][num2]int) (sumaDeElementos, cantElementos int) {
	for f := 0; f < num1; f++ {
		for c := 0; c < num2; c++ {
			if a[f][c] >= 5 {
				sumaDeElementos += a[f][c]
				cantElementos++
			}
		}
	}
	return sumaDeElementos, cantElementos
}

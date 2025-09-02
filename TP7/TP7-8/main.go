package main

import "fmt"

const (
	num1 int = 3
	num2 int = 5
)

func main() {
	var (
		a     [num1][num2]int
		mayor int
	)
	for f := 0; f < num1; f++ {
		for c := 0; c < num2; c++ {
			fmt.Print("Ingrese los numeros:")
			fmt.Scanln(&a[f][c])
		}
	}

	for f := 0; f < num1; f++ {
		for c := 0; c < num2; c++ {
			if f == 0 {
				mayor = a[f][c]
			} else if mayor < a[f][c] {
				mayor = a[f][c]
			}
		}
	}

	fmt.Print("El numero mas grande de los ingresados es ", mayor)
}

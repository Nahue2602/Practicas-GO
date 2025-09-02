package main

import "fmt"

const num1 int = 3

func main() {
	var (
		a  [num1][num1]int
		cf int
		cc int
	)

	for f := 0; f < num1; f++ {
		for c := 0; c < num1; c++ {
			fmt.Print("Ingrese los numeros:")
			fmt.Scanln(&a[f][c])
		}
	}

	for f := 0; f < num1; f++ {
		for c := 0; c < num1; c++ {
			if f+1 == a[f][c] {
				cf++
			}
			if c+1 == a[f][c] {
				cc++
			}
		}
	}

	fmt.Printf("La cantidad de elementos que coinciden con sus filas y columnas es iguar a %v filas y %v columnas", cf, cc)
}

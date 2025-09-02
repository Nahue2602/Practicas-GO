package main

import "fmt"

const num1 int = 10

func main() {
	var a [num1][num1]int

	for f := 0; f < num1; f++ {
		for c := 0; c < num1; c++ {
			fmt.Print("Ingrese los numeros:")
			fmt.Scanln(&a[f][c])
		}
	}
	sumP, sumIm := sumaParesImpares(a)
	fmt.Println("La suma de los numeros pares es de", sumP)
	fmt.Println("La suma de los numeros impares es de", sumIm)
}

func sumaParesImpares(a [num1][num1]int) (sumP, sumIm int) {
	for f := 0; f < num1; f++ {
		for c := 0; c < num1; c++ {
			if a[f][c]%2 == 0 {
				sumP += a[f][c]
			} else {
				sumIm += a[f][c]
			}
		}
	}

	return sumP, sumIm
}

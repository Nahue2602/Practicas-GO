package main

import "fmt"

const num1 int = 4

func main() {
	var a [num1][num1]int

	for f := 0; f < num1; f++ {
		for c := 0; c < num1; c++ {
			fmt.Print("Ingrese un numero:")
			fmt.Scanln(&a[f][c])
		}
	}
	resultado := suma(a)
	fmt.Println("la suma de todos los numeros es de ", resultado)
}

func suma(a [num1][num1]int) (result int) {
	for f := 0; f < num1; f++ {
		for c := 0; c < num1; c++ {
			result += a[f][c]
		}
	}
	return result
}

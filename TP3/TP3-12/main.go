package main

import (
	"fmt"
)

func main() {

	const (
		count = 10
	)

	var (
		cantMul int
		num     int
	)

	for i := 0; i < count; i++ {

		fmt.Print("ingrese un numero:")
		fmt.Scanln(&num)

		if num%5 == 0 {
			cantMul++
		}
	}

	fmt.Print("la cantidad de numeros multiplos de 5 ingresados es de ", cantMul)
}

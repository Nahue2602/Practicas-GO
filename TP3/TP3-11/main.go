package main

import (
	"fmt"
)

func main() {

	const (
		count = 10
	)

	var (
		sumIm int
		sumPa int
		num   int
	)

	for i := 0; i < count; i++ {
		fmt.Print("ingrese un numero:")
		fmt.Scanln(&num)

		if num%2 == 0 {
			sumPa += num
		} else {
			sumIm += num
		}
	}

	fmt.Printf("La suma de los numeros pares es %v y la suma de los numeros impares es %v", sumPa, sumIm)

}

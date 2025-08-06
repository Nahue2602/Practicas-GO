package main

import (
	"fmt"
)

func main() {

	const (
		count = 10
	)

	var (
		mayor int
		menor int
		num   int
	)

	for i := 0; i < count; i++ {

		fmt.Print("ingrese un numero:")
		fmt.Scanln(&num)

		if i == 0 {
			mayor = num
			menor = num
		}

		if num < menor {
			menor = num
		} else if num > mayor {
			mayor = num
		}
	}

	fmt.Printf("el numero mas grande es %v y el numero mas chico es %v", mayor, menor)
}

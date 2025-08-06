package main

import (
	"fmt"
)

func main() {

	const fin = 8

	var (
		num      int
		suma     int
		promedio float64
	)

	for i := 0; i < fin; i++ {
		fmt.Print("Ingresa un numero:")
		fmt.Scanln(&num)
		suma += num
	}

	promedio = float64(suma) / float64(fin)

	fmt.Print("El promedio de esos sumeros es de:", promedio)

}

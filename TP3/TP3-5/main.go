package main

import (
	"fmt"
	"math"
)

func main() {

	var num int

	fmt.Println("ingrese un numero entero")
	fmt.Scanln(&num)
	v := math.Pow(-1, float64(num))
	if v == 1 {
		fmt.Println("Su numero es par")
	} else {
		fmt.Println("Su numero es impar")
	}
}

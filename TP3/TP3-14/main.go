package main

import (
	"fmt"
)

func main() {

	var (
		num   int
		count int
		sum   int
		auxi  int
	)

	fmt.Print("ingrese un numero:")
	fmt.Scanln(&num)
	count = num
	auxi = num
	for i := 0; i < count; i++ {
		sum += auxi
		auxi = auxi - 1
	}

	fmt.Print("la suma de ese numero y todos sus anteriores es ", sum)
}

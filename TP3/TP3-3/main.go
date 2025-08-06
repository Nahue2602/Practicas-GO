package main

import "fmt"

func main() {
	var (
		n1 int
		n2 int
	)

	fmt.Println("ingrese 2 numeros")
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	if n1 > n2 {
		fmt.Println(n1, " es el mayor")
	} else {
		fmt.Println(n2, " es el mayor")
	}
}

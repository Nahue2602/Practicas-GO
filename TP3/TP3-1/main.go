package main

import "fmt"

func main() {

	var (
		n1  int
		n2  int
		n3  int
		dif int
		pro int
		div int
	)

	fmt.Println("Ingrese 3 numeros enteros: ")
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Scanln(&n3)
	dif = n1 - n2
	pro = n2 * n3
	div = n1 / n3
	fmt.Printf("la diferencia de los dos primeros es %v, el producto de los dos ultimos es %v y la division del primero y el tercero es %v", dif, pro, div)
}

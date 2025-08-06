package main

import "fmt"

const (
	PI = 3.141516
	t  = 2
)

func main() {

	var (
		r float64
		a float64
		l float64
	)

	fmt.Println("Ingrese el radio: ")
	fmt.Scanln(&r)
	a = PI * r * r
	l = t * PI * r
	fmt.Printf("La longitud es de %v y su area es de %v", l, a)

}

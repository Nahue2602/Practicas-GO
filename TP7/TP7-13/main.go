package main

import "fmt"

const (
	num1 int = 10
	num2 int = 20
)

func main()  {
	var (
		a[num1][num2] int
		veri1 bool
		veri2 bool
	)

	for f := 0; f < num1; f++ {
		for c := 0; c < num2; c++ {
			fmt.Printf("Ingrese el elemento de la fila %v y la columna %v:", f+1, c+1)
			fmt.Scanln(&a[f][c])
		}
	}
	fmt.Println("¿Desea verificar si es mayor?")
	fmt.Scanln(&veri1)
	fmt.Println("¿Desea verificar en las filas pares?")
	fmt.Scanln(&veri2)

	n1, n2 := verificar(a,veri1,veri2)
}

func verificar(a[num1][num2]int, veri1, veri2 bool) (n1, n2 int) {
	
}
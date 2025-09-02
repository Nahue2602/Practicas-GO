/*
Se tienen 2 vectores A y B de 20 elementos. Programar un procedimiento que permita ordenar
vectores de forma ascendente. El procedimiento debe recibir como parámetro por referencia el
vector a ordenar.
Una vez ordenado los 2 vectores mediante la llamada al procedimiento, generar un nuevo vector C
que sea la intercalación ordenada de A y de B (considere que no hay elementos repetidos).
*/
package main

import "fmt"

const (
	num1 = 20
	num2 = 19
	num3 = 40
)

func main() {
	var (
		a [num1]int
		b [num1]int
		c [num3]int
	)
	for i := 0; i < num1; i++ {
		fmt.Print("Ingrese un numero para el vector A:")
		fmt.Scanln(&a[i])
	}

	for i := 0; i < num1; i++ {
		fmt.Print("Ingrese un numero para el vector B:")
		fmt.Scanln(&b[i])
	}
	ordenarVector(&a)
	ordenarVector(&b)
	cargarVectorC(a, b, &c)
	fmt.Println("El vector A es:", a)
	fmt.Println("El vector B es:", b)
	fmt.Println("El vector C es:", c)
}

func ordenarVector(v *[num1]int) {
	for x := 0; x < num2; x++ {
		for y := x + 1; y < num1; y++ {
			if v[x] > v[y] {
				aux := v[x]
				v[x] = v[y]
				v[y] = aux
			}
		}
	}
}

func cargarVectorC(a [num1]int, b [num1]int, c *[num3]int) {
	iA, iB, iC := 0, 0, 0

	for iA < num1 && iB < num1 {
		if a[iA] < b[iB] {
			c[iC] = a[iA]
			iA++
			iC++
		} else if a[iA] > b[iB] {
			c[iC] = b[iB]
			iB++
			iC++
		} else {
			c[iC] = a[iA]
			iA++
			iB++
			iC++
		}
	}

	for iA < num1 {
		c[iC] = a[iA]
		iA++
		iC++
	}

	for iB < num1 {
		c[iC] = b[iB]
		iB++
		iC++
	}
}

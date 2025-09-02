package main

import "fmt"

const num1 int = 5

func main() {
	var (
		a   [num1][num1]int
		num int
	)

	for f := 0; f < num1; f++ {
		for c := 0; c < num1; c++ {
			fmt.Print("Ingrese los numeros:")
			fmt.Scanln(&a[f][c])
		}
	}

	fmt.Print("Ingrese el numero que quiere verificar si aparece en la matriz y cuantas veces esta en la misma:")
	fmt.Scanln(&num)

	resultado := verificarNumero(a, num)
	if resultado != 0 {
		fmt.Printf("El numero ingresado anteriormente aparece %v veces", resultado)
	} else {
		fmt.Print("El numero ingresado anteriormente no aparece ninguna vez")
	}
}
func verificarNumero(a [num1][num1]int, num int) (contador int) {
	for f := 0; f < num1; f++ {
		for c := 0; c < num1; c++ {
			if a[f][c] == num {
				contador++
			}
		}
	}
	return contador
}

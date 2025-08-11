package main

import (
	"fmt"
)

const count = 5

func main() {
	var (
		num  int
		num2 int
	)

	for i := 0; i < count; i++ {
		fmt.Print("Ingrese el numero a evaluar:")
		fmt.Scanln(&num)
		num2 = num
		verificarMultiplo(num, num2)

	}

}

func verificarMultiplo(num, num2 int) (resultado bool) {
	for num >= 3 {
		num -= 3
	}

	for num2 >= 5 {
		num2 -= 5
	}

	if num == 0 && num2 != 0 {
		resultado = true
		fmt.Println("El numero cumple con la condicion")
	} else {
		resultado = false
		fmt.Println("El numero no cumple con la condicion")
	}

	return resultado
}

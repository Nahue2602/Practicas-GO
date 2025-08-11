package main

import "fmt"

func main() {
	var (
		a      float64
		b      float64
		accion int = 1
	)
	fmt.Print("Ingrese el primer numero:")
	fmt.Scanln(&a)
	fmt.Println("")
	fmt.Print("Ingrese el segundo numero:")
	fmt.Scanln(&b)
	fmt.Println("")
	for accion != 0 {
		fmt.Println("¿Que operacion desea realizar con dichos numeros?")
		fmt.Println("1: Suma")
		fmt.Println("2: Resta")
		fmt.Println("3: Multiplicar")
		fmt.Println("4: Dividir")
		fmt.Println("0: Salir")
		fmt.Scanln(&accion)
		fmt.Println("")
		switch accion {
		case 1:
			fmt.Println("eligio SUMA")
			resultado := suma(a, b)
			fmt.Println("La suma de los numeros es", resultado)
			fmt.Println("")
		case 2:
			fmt.Println("eligio RESTA")
			resultado := resta(a, b)
			fmt.Println("La resta de los numeros es", resultado)
			fmt.Println("")
		case 3:
			fmt.Println("eligio MULTIPLICACION")
			resultado := multiplicacion(a, b)
			fmt.Println("El producto de los numeros es", resultado)
			fmt.Println("")
		case 4:
			fmt.Println("eligio DIVISION")
			resultado := division(a, b)
			fmt.Println("La division de los numeros es", resultado)
			fmt.Println("")
		}
	}
	fmt.Print("Eligio SALIR")
	fmt.Println("")
}
func suma(a, b float64) (resultado float64) {
	return a + b
}
func resta(a, b float64) (resultado float64) {
	return a - b
}
func multiplicacion(a, b float64) (resultado float64) {
	return a * b
}
func division(a, b float64) (resultado float64) {
	return a / b
}

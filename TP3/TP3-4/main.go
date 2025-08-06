package main

import "fmt"

func main() {

	var (
		num int
	)

	fmt.Println("ingresar un numero del 1 al 7")
	fmt.Scanln(&num)

	switch num {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Este numero no corresponde a la consigna asignada")
	}
}

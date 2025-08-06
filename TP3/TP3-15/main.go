package main

import (
	"fmt"
)

func main() {

	const f = 10

	var (
		nom     string
		edad    int
		men     int
		entre   int
		pro     float64
		nomMen  string
		nomMay  string
		edadMen int
		edadMay int
		suma    int
	)

	for i := 0; i < f; i++ {
		fmt.Print("ingresa el nombre del alumno:")
		fmt.Scanln(&nom)
		fmt.Print("ingresa la edad del alumno:")
		fmt.Scanln(&edad)

		suma += edad

		if i == 0 {
			nomMen = nom
			edadMen = edad
			nomMay = nom
			edadMay = edad
		}

		if edadMen > edad {
			nomMen = nom
			edadMen = edad
		}

		if edadMay < edad {
			nomMay = nom
			edadMay = edad
		}

		if edad < 12 {
			men++
		}

		if edad > 10 && edad < 16 {
			entre++
		}
		pro = float64(suma) / f
	}

	fmt.Println("el alumno mas joven es: ", nomMen)
	fmt.Println("el alumno mas grande es: ", nomMay)
	fmt.Println("la cantidad de alumnos menores a 12 años es de: ", men)
	fmt.Println("la cantidad dee alumnos menores que 16 pero mayores a 10 es de: ", entre)
	fmt.Println("el promedio de todas las edades es de: ", pro)
}

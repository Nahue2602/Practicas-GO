package main

import "fmt"

func main() {

	var (
		canAlu     int
		canAluRegu int
		por        float64
		nMat       int
		n1         int
		n2         int
		n3         int
	)

	fmt.Print("Ingrese en numero de matricula del alumno:")
	fmt.Scanln(&nMat)

	for nMat != 0 {
		cantParcial := 0
		canAlu++
		fmt.Print("Ingrese la nota del primer parciarl:")
		fmt.Scanln(&n1)
		fmt.Print("Ingrese la nota del segundo parciarl:")
		fmt.Scanln(&n2)
		fmt.Print("Ingrese la nota del tercer parciarl:")
		fmt.Scanln(&n3)

		if n1 >= 7 {
			cantParcial++
		}
		if n2 >= 7 {
			cantParcial++
		}
		if n3 >= 7 {
			cantParcial++
		}

		if cantParcial >= 2 {
			canAluRegu++
		}
		fmt.Print("Ingrese en numero de matricula del alumno:")
		fmt.Scanln(&nMat)
	}
	por = (float64(canAluRegu) * 100) / float64(canAlu)
	fmt.Println("La cantidad total de alumnos es de", canAlu)
	fmt.Printf("La cantidad de alumnos regulares es de %v y el porcenaje de estos sobre el total es de %v", canAluRegu, por)
}

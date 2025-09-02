package main

import "fmt"

const (
	mat int = 4
	alu int = 2
)

func main() {
	var (
		AL          [alu][mat]float64
		aluExaminar int
	)

	for a := 0; a < alu; a++ {
		for m := 0; m < mat; m++ {
			fmt.Printf("Ingrese la nota del alumno %d, en la materia %d:", a+1, m+1)
			fmt.Scanln(&AL[a][m])
		}
	}

	promedioPorMateria := calcularPromedioMateria(AL)
	promedioPorEstudiante := calcularPromedioEstudiante(AL)
	fmt.Print("Ingrese el numero de estudiante del cual desea saber cuantas materias aprobó y cuantas desaprobó:")
	fmt.Scanln(&aluExaminar)
	cA, cD := materiasAprobadasYDesaprobadas(AL, aluExaminar)

	fmt.Println("El promedio de cada meteria es de", promedioPorMateria)
	fmt.Println("El promedio de cada estudiante es de", promedioPorEstudiante)
	fmt.Printf("el alumno %v tiene %v materias aprobadas y %v materias desaprobadas", aluExaminar, cA, cD)

}

func calcularPromedioMateria(AL [alu][mat]float64) (V [mat]float64) {
	var (
		suma float64
	)
	for m := 0; m < mat; m++ {
		for a := 0; a < alu; a++ {
			suma += AL[a][m]
		}
		promedio := suma / float64(alu)
		suma = 0
		V[m] = promedio
	}

	return V
}

func calcularPromedioEstudiante(AL [alu][mat]float64) (V [alu]float64) {
	var (
		suma float64
	)

	for a := 0; a < alu; a++ {
		for m := 0; m < mat; m++ {
			suma += AL[a][m]
		}
		promedio := suma / float64(mat)
		suma = 0
		V[a] = promedio
	}

	return V
}

func materiasAprobadasYDesaprobadas(AL [alu][mat]float64, num int) (cA, cD int) {
	for m := 0; m < mat; m++ {
		if AL[num-1][m] >= 6 {
			cA++
		} else {
			cD++
		}
	}

	return cA, cD
}

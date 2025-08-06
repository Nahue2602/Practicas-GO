package main

import (
	"fmt"
)

func main() {

	var (
		nom              string
		edad             int
		salario          int
		sumTot           int
		empleadosMenores int
		cantEmpleados    int
		nomEmpleado      string
	)

	fmt.Print("ingrese el la edad del empleado:")
	fmt.Scanln(&edad)

	for edad != 0 {

		fmt.Print("ingrese en nombre del empleado:")
		fmt.Scanln(&nom)
		fmt.Print("ingrese en salario del empleado:")
		fmt.Scanln(&salario)

		sumTot += salario
		cantEmpleados++

		if edad > 25 && salario > 3000 {
			nomEmpleado += nom + ","
		} else if edad < 18 {
			empleadosMenores++
		}
		fmt.Print("ingrese el la edad del empleado:")
		fmt.Scanln(&edad)
	}
	fmt.Println("los empleados mayores de 25 y con sueldos mayores a $3000 se llaman:", nomEmpleado)
	fmt.Println("la cantidad de empleados menores de 18 es de:", empleadosMenores)
	fmt.Println("la cantidad de empleados ingresados es de:", cantEmpleados)
	fmt.Println("la suma de los sueldos es de:", sumTot)

}

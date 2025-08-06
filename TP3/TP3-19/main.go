/*
Una empresa de viajes que realizó una campaña de promoción necesita calcular las comisiones que deberá
pagar a sus promotores de ventas. Las mismas se calculan según la cantidad de excursiones vendidas por
cada uno de ellos. Se decide asignar una categoría a cada vendedor según la cantidad vendida y según esa
categoría se pagará un importe por cada excursión de acuerdo a la siguiente tabla:

Categoría Excursiones $ x Excursión
A <10 $1000
B >= 10 y < 50 $1500
C >= 50 y < 100 $1700
D >= 100 $1900

Se cuenta con N datos que corresponden a las cantidades vendidas por cada promotor. Se necesita
conocer:
• El importe a pagarle a cada promotor
• La cantidad de promotores de cada categoría
*/

package main

import (
	"fmt"
)

func main() {

	var (
		nom   string
		n     int
		cantA int
		cantB int
		cantC int
		cantD int
	)

	fmt.Print("ingrese la cantidad de viajes vendidos: ")
	fmt.Scanln(&n)

	for n != 0 {
		fmt.Print("ingrese el nombre del vendedor: ")
		fmt.Scanln(&nom)

		if n < 10 {
			fmt.Println("El importe a pagarle a", nom, "es de $1000")
			cantA += 1
		} else if n >= 10 && n < 50 {
			fmt.Println("El importe a pagarle a", nom, "es de $1500")
			cantB += 1
		} else if n >= 50 && n < 100 {
			fmt.Println("El importe a pagarle a", nom, "es de $1700")
			cantC += 1
		} else if n >= 100 {
			fmt.Println("El importe a pagarle a", nom, "es de $1900")
			cantD += 1
		}
		fmt.Print("ingrese la cantidad de viajes vendidos: ")
		fmt.Scanln(&n)
	}
	fmt.Println("La cantidad de vendedores de clase A es de ", cantA)
	fmt.Println("La cantidad de vendedores de clase B es de ", cantB)
	fmt.Println("La cantidad de vendedores de clase C es de ", cantC)
	fmt.Println("La cantidad de vendedores de clase D es de ", cantD)
}

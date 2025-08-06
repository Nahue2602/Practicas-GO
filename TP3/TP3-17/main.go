package main

import (
	"fmt"
)

func main() {

	var (
		num           int
		muj           int
		varDelTerncer int
		cUno          int
		cDos          int
		cTres         int
		sexo          string
		voto          int
	)

	fmt.Print("Ingrese la cantidad de votantes:")
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		fmt.Print("ingrese el sexo (m o h):")
		fmt.Scanln(&sexo)
		fmt.Print("ingrese el candidato (1, 2 o 3):")
		fmt.Scanln(&voto)

		if voto == 1 {
			cUno++
		} else if voto == 2 {
			cDos++
		} else {
			cTres++
		}

		if sexo == "m" {
			muj++
		} else if voto == 3 {
			varDelTerncer++
		}
	}

	fmt.Println("la cantidad de mujeres que votaron es de", muj)
	fmt.Println("la cantidad de varones que votaron al candidato 3 es de", varDelTerncer)

	if cUno > cDos && cUno > cTres {
		fmt.Println("El candidato 1 es el ganador")
	} else if cDos > cUno && cDos > cTres {
		fmt.Println("El candidato 2 es el ganador")
	} else {
		fmt.Println("El candidato 3 es el ganador")
	}

}

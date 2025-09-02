package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		num, random int
		intentos    = 1
	)
	rand.Seed(time.Now().Unix())
	random = rand.Intn(100)
	fmt.Print("Ingrese un numero:")
	fmt.Scanln(&num)
	for {

		numeroAdivinado := checkNumeroAdivinado(random, num)

		checkIntento(numeroAdivinado, intentos)

		if numeroAdivinado == 0 {
			break
		}
		intentos++
		fmt.Print("Ingrese un numero:")
		fmt.Scanln(&num)
	}
}

func checkNumeroAdivinado(random, input int) int {
	if input == random {
		return 0
	} else if input > random {
		return 1
	} else {
		return -1
	}

}

func checkIntento(check, intentos int) {

	switch check {
	case 0:
		fmt.Println("El numero es correcto")
	case 1:
		fmt.Println("El mumero es mayor")
	default:
		fmt.Println("El numero es menor")
	}

	if intentos <= 5 {
		fmt.Println("Bueno")
	} else if intentos > 15 {
		fmt.Println("malo")
	} else {
		fmt.Println("Regular")
	}

}

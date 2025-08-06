package main

import "fmt"

func main() {

	var (
		h string
		m string
		s string
	)

	fmt.Print("ingrese la hora:")
	fmt.Scanln(&h)
	fmt.Print("ingrese los minutos:")
	fmt.Scanln(&m)
	fmt.Print("ingrese los segundos:")
	fmt.Scanln(&s)
	fmt.Printf("el horario completo es: %v:%v:%v", h, m, s)
}

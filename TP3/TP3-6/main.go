package main

import "fmt"

func main() {

	var (
		n1 int
		n2 int
		n3 int
	)
	fmt.Println("Ingrese 3 numeros entros:")
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Scanln(&n3)

	if n1 > n2 {
		if n1 > n3 {
			fmt.Println(n1, " Es el mayor de los 3")
		}
	} else if n2 > n3 {
		fmt.Println(n2, " Es el mayor de los 3")
	} else {
		fmt.Println(n3, " Es el mayor de los 3")
	}
}

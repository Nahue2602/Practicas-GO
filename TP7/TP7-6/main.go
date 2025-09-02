package main

import "fmt"

const num1 int = 4

func main() {
	var (
		a [num1][num1]int
	)
	for f := 0; f < num1; f++ {
		for c := 0; c < num1; c++ {
			fmt.Print("Ingrese los numeros")
			fmt.Scanln(&a[f][c])
		}
	}

	for f := 0; f < num1; f++ {
		for c := 0; c < num1; c++ {
			if (f+1)%2 == 0 {
				fmt.Print(a[f][c], " ")
			}
		}
		fmt.Println("")
	}
}

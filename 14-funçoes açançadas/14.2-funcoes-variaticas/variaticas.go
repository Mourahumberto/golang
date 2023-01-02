package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println(numeros)
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	totalSoma := soma(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(totalSoma)
}

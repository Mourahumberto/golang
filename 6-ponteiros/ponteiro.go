package main

import "fmt"

//ponteiro é uma referência de memória

func main() {
	var var1 int
	var ponteiro1 *int

	var1 = 100
	ponteiro1 = &var1

	fmt.Println(var1, ponteiro1)
	fmt.Println(var1, *ponteiro1)
}

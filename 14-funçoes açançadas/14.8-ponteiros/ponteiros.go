package main

import "fmt"

func dobra(numero int) int {
	return numero * 2
}

func dobraPonteiro(numero *int) {
	*numero = *numero * 2
}
func main() {
	numero := 20
	numeroDobrado := dobra(numero)
	fmt.Println(numeroDobrado)
	fmt.Println(numero)

	novoNumero := 20
	fmt.Println(novoNumero)
	dobraPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}

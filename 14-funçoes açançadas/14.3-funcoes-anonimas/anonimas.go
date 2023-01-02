package main

import "fmt"

func main() {
	func(texto string) {
		fmt.Println(texto)
	}("testes de funções")

	retorno := func(texto string) string {
		return fmt.Sprintf(texto)
	}("Com retornos")
	fmt.Println(retorno)
}

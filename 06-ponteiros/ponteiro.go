package main

import "fmt"

//ponteiro é uma referência de memória

func main() {
	var var1 int       // salva a variável
	var ponteiro1 *int // salva o endereço de memória

	var1 = 100
	ponteiro1 = &var1 // para salvar o endereço de memória é necessário esse &

	fmt.Println(var1, ponteiro1)
	fmt.Println(var1, *ponteiro1) // Para mostrar o valor dentro de um endereço de memória é necessário o *
}

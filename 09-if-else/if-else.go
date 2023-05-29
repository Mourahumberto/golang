package main

import "fmt"

func main() {
	numero := 8

	if numero > 15 {
		fmt.Println("é maior que 15")
	} else {
		fmt.Println("é menor ou igual a 15")
	}

	// esse é o if inite tem o objetivo de criar uma variável onde o escopo dela pode ser acessado apenas dentro da estrutura.
	if outroNumero := numero; outroNumero > 10 {
		fmt.Println("maior que 10")
	} else if numero == 10 {
		fmt.Println("é igual a 10")
	} else {
		fmt.Println("menor ou igual a 10")
	}
	fmt.Println(numero)
	//fmt.Println(outroNumero) // daria erro pois ela só é conhecida dentro do escopo do if e else onde foi criada.
}

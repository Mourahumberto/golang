package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("neto@gmail.com")
	fmt.Println(erro)
}

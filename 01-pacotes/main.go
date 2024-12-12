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
	if erro == nil {
		fmt.Println("Sem error")
	}

}

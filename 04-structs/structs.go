package main

import "fmt"

// Proximo do que seria uma class

type usuario struct {
	nome  string
	idade int
}

type usuario2 struct {
	nome    string
	idade   int
	enderec endereco
}

type endereco struct {
	rua  string
	casa int
}

func main() {
	fmt.Println()
	//forma1
	var u1 usuario

	u1.idade = 21
	fmt.Println(u1)

	//forma2
	u2 := usuario{"Davi", 21}
	fmt.Println(u2)

	//forma3
	u3 := usuario{idade: 22}
	u4 := usuario{nome: "Davi"}

	fmt.Println(u3)
	fmt.Println(u4)

	//struct in struct

	var u5 usuario2
	u5.nome = "Humberto"
	u5.idade = 28
	u5.enderec.rua = " WL12"
	u5.enderec.casa = 11

	fmt.Println(u5)
}

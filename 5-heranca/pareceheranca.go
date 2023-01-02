package main

import "fmt"

type usuario2 struct {
	nome  string
	idade int
	endereco
}

type endereco struct {
	rua  string
	casa int
}

func main() {

	//struct in struct

	var u5 usuario2
	u5.nome = "Humberto"
	u5.idade = 28
	u5.rua = " WL12"
	u5.casa = 11
	fmt.Println(u5)

}

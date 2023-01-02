package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("salvando os dados do usuário %s no banco \n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}
func main() {
	usuario1 := usuario{"Neto", 28}
	usuario1.salvar()
	fmt.Println(usuario1.maiorDeIdade())
	usuario2 := usuario{"May", 28}
	usuario2.salvar()

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

}

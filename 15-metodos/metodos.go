package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// isso é como criamos um método para nossa struc. nesse caso quer dizer que todos os usuários tem o método salvar()
// func (u usuario) salvar() {}

func (u usuario) salvar() {
	fmt.Printf("salvando os dados do usuário %s no banco \n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// precisamos passar o ponteiro, se não ele cria uma cópia e não salva o valor do que foi passado.
// Então se for preciso alterar um valor é necessário um ponteiro.
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

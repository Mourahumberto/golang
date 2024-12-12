package main

import "fmt"

func funcao1() {
	fmt.Println("executando a função1")
}

func funcao2() {
	fmt.Println("executando a função2")
}

func funcao3() {
	fmt.Println("executando a função3")
}

func alunoAprovado(n1, n2 int) bool {
	defer fmt.Println("Calculando a média do aluno")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	funcao1()
	// tem a função de adiar a função até o ultimo momento.
	// ele adia a função, mas ele executa antes do return
	defer funcao2()
	funcao3()
	alunoAprovado(7, 8)
}

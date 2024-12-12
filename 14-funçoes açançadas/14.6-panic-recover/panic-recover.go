package main

import "fmt"

func recuperaFunc() {
	if r := recover(); r != nil {
		fmt.Println("Função recuperada")
	}
}
func alunoAprovado(n1, n2 float64) bool {
	defer recuperaFunc()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	}
	if media < 6 {
		return false
	}
	// Quando o panic é executado ele para a execução do programa.
	// mas antes dele parar o programa ele executa todos os defers
	panic(" A media é exatamente 6 ")
}

func main() {
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pós execução")
}

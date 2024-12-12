package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println(" Incrementando i", i)
		// time.Sleep(time.Second)

	}
	fmt.Println("saiu")
	fmt.Println(i)

	// assim como o if init o j é uma variável que só existe dento desse escopo
	for j := 0; j < 10; j++ {
		fmt.Println(" Incrementando j", j)
		// time.Sleep(time.Second)

	}
	fmt.Println("saiu do j")
	// fmt.Println(j) daria erro não existe o j fora daquele escopo.

	nomes := [3]string{"joão", "neto", "kaua"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
	for _, nome := range nomes {
		fmt.Println(nome)
	}
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for indice, num := range slice1 {
		fmt.Println(indice, string(num))
	}

	usuario := map[string]string{
		"nome":      "neto",
		"sobrenome": "moura",
	}

	for chave, valor := range usuario {
		fmt.Println("chave:", chave, "valor:", valor)
	}
	// looop infinito
	for {
		fmt.Println("loop infinito")
		time.Sleep(time.Second)
	}
}

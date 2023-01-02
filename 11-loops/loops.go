package main

import (
	"fmt"
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

}

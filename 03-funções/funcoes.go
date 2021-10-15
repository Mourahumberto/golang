package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculos(n1, n2 int) (int, int) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	soma := somar(30, 10)
	println(soma)

	retornasoma, retornasub := calculos(30, 20)
	fmt.Println(retornasoma, retornasub)

	retornasoma1, _ := calculos(30, 20)
	fmt.Println(retornasoma1)
}

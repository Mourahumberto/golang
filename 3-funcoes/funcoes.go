package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculo(n1 int, n2 int) (int, int) {
	sub := n2 - n1
	som := n1 + n2
	return som, sub

}
func calculo2() (int, int) {
	sub := 10 - 5
	som := 10 + 5
	return sub, som
}

func main() {
	soma := somar(10, 20)
	cal1, cal2 := calculo(10, 20)
	_, test2 := calculo(50, 30)
	test1, _ := calculo(60, 20)
	fmt.Println(soma)
	fmt.Println(cal1, cal2)
	fmt.Println(test1)
	fmt.Println(test2)
	teste1, teste2 := calculo2()
	fmt.Println(teste1, teste2)
}

package main

import (
	"errors"
	"fmt"
)

var varglobal string = "variavel global"

func main() {
	var var1 string = "variavel1"
	var2 := "variavel2"
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(varglobal)

	var (
		var3 string = "variavel3"
		var4 string = "variavel4"
	)
	fmt.Println(var3)
	fmt.Println(var4)

	const const1 string = "constante1"
	fmt.Println(const1)

	var inteiros1 int = 12
	println(inteiros1)
	inteiros2 := 13
	println(inteiros2)

	var float1 float64 = 12300.1
	println(float1)
	float2 := 1300000.3
	println(float2)

	var boleano1 bool = false
	fmt.Println(boleano1)

	var errooou error = errors.New("Erro interno")
	fmt.Println(errooou)

}

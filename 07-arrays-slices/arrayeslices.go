package main

import "fmt"

func main() {
	// Array só pode ter dados do mesmo tipo e de tamnhos definidos na variável
	var array1 [3]string
	fmt.Println(array1)
	array1[0] = "teste1"
	array1[2] = "teste2"
	//array1[3] = "teste3" daria erro pois o array só tem 3 posições
	fmt.Println(array1)

	array2 := [4]string{"teste1", "test2", "test3", "test4"}
	fmt.Println(array2)

	// Slice só pode ter dados do mesmo tipo. Poém podemos usar um tamanho variável

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// var slice2 []int
	slice2 := append(slice1, 9)

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1[3] = 9
	fmt.Println(slice1)

	// pegando partes do slice
	slice3 := slice1[3:6]
	fmt.Println(slice3)
}

package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2)
	canal <- "Olá mundo1!"
	canal <- "Olá mundo2!"
	// canal <- "Olá mundo3!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}

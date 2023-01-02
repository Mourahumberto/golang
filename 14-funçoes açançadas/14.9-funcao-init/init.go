package main

import "fmt"

func main() {
	fmt.Println("funçao main sendo executada.")
}

func init() {
	fmt.Println("uma função que acontece antes do main, serve geralmente pra fazer o set up de algo antes do código")
}

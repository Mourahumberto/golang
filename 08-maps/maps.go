package main

import "fmt"

// Parecido com struct porém só aceita um tipo de dado por map
func main() {

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "silva",
		//"idade":	 23 // não daria certo pois a idade estã como int e o map só aceita tudo string ou tudo int ou outro tipo, porém tudo o mesmo.
	}
	fmt.Println(usuario["nome"])
}

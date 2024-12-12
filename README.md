# golang

## Introdução GO

## 1 - Explicando hello World

- arquivo main.go
```go
package main

import "fmt"

func main() {
    // Inicio
	fmt.Println("Hello World!")
    // Fim
}
```
### 1.1 - explicando o código
- Todo Projeto Go deve ter o package main, que será o pacote principal, que será executado no run
- O pacote main deve ter uma função main, que será a função que irá iniciar e finalizar a aplicação
- import "fmt" : importará o println
- rodando a aplicação: ``` go run main.go```

## 2 - Criando modulos

- Pela linha de comando cria um módulo.
- ```go mod init modulo```

### 2.1 - Explicando sobre Pacotes
- Exemplo: 01-pacotes

- irá pegar os módulos com o nome package auxiliar
```go
import (
    "fmt"
    "modulo/auxiliar"
)
} 
```

- Irá referenciar o pacote.Classe. a classe Escrever está com letra maiuscula. Pois dessa forma ele pode ser acessado fora do pacote. 
```go
	auxiliar.Escrever()
```

- diferente do arquivo auxiliar/auxiliar2.go a classe está pacote.classe, Desta forma ela só pode ser acessada dentro do pacote. Por isso a classe do escrever2 só pode ser acessado pelo escrever 1.
```go
	escrever2()
```

### 2.2 baixando novosw pacotes para importar
$ go get https://github.com/badoux/checkmail
- e depois importa no código
- se os modulos não são referenciado só dar um "go mod tidy" que irá tirar os módulos não usados do go.mod

## 3 - Variáveis
- Exemplo: 02-variaveis

- Formas de declaração de variáveis.
- variáveis que são declarados e não são usadas dão erro no código.

- Podem ser vistas no scopo de todo o package
```go
var varglobal string = "variavel global"
```

- Só podem ser vistos dentro da classe
```go
var var1 string = "variavel1"
```
- declarando com o gopher, operador de declaração rápida. Ele já atribui o tipo dinamicamente.
- só pode ser usado esse operador dentro da classe
```go
var2 := "variavel2"
```

### 3.1 - tipos
- Exemplo: 03-funções

- string: letras
- int: numeros inteiros
- float64 e float32: números de ponto flutuante
- bool: true ou false
- error: erros em go são um tipo.

## 4 - Funções
- As funções no go podem ter um ou mais retornos, diferente de outras linguagens de programação

- Nesta função verificamos, que a função precisa de duas entradas do tipo inteiro, e retorna uma saída do tipo inteiro.
```go
func somar(n1 int, n2 int) int {
	return n1 + n2
}
```

- Nesta função precisa de duas entradas e retorna duas saídas. 
```go
func calculos(n1, n2 int) (int, int) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}
```

- Função main chamando as funções calculos, verifique que nela colocamos a entrada. Caso só precise de uma saída defina a outra como underline, pois dessa não dará erro por não usar uma variável definida no código.
```go
func main() {
	soma := somar(30, 10)
	println(soma)

	retornasoma, retornasub := calculos(30, 20)
	fmt.Println(retornasoma, retornasub)

	retornasoma1, _ := calculos(30, 20)
	fmt.Println(retornasoma1)
}
```


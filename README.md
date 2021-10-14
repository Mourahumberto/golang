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

### 2.1 - Explicando o 01-modilos

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

## 3 - Variáveis


package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// if comando == 1 { //o if sempre necessita que o retorno seja true ou false
	// 	fmt.Println("comando 1")
	// } else if comando == 2 {
	// 	fmt.Println("comando 2")
	// } else if comando == 0 {
	// 	fmt.Println("comando 0")
	// } else {
	// 	fmt.Println("Não conheço essa opção")
	// }
	var nome string = "Neto"
	var versao float32 = 0.1
	fmt.Println("Olá senhor", nome, "estamos na versão", versao)

	for {
		exibeIntro()
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoracao()
		case 2:
			mostrarLogs()
		case 0:
			fmt.Println("Sair do programa")
			os.Exit(0)
		default:
			fmt.Println("não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntro() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")
}

func leComando() int {
	var leComando int
	fmt.Scan(&leComando) // comando Scan espera algo nesse caso um inteiro
	fmt.Println(leComando)
	return leComando

}

func iniciarMonitoracao() {
	fmt.Println("Monitorando")
	sites := leSitesDoArquivo()
	for i := 0; i < 2; i++ {
		for _, site := range sites { // comando range faz uma leitura em todos os valores do array
			resp, _ := http.Get(site)
			if resp.StatusCode == 200 {
				fmt.Println("Site:", site, "foi carregado com sucesso")
				registraLog(site, true)
			} else {
				fmt.Println("Site:", site, "teve problema ao carregar.")
				registraLog(site, false)
			}
		}
		time.Sleep(2 * time.Second)
	}
}

func leSitesDoArquivo() []string {
	// arquivo, erro := ioutil.ReadFile("sites.txt") // esse pacote ReadFile le um arquivo e transforma em binário.
	var sites []string
	arquivo, erro := os.Open("sites.txt")
	if erro != nil {
		fmt.Println("Ocorreu um erro:", erro)

	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, erro := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if erro == io.EOF {
			break
		}

	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, erro := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if erro != nil {
		fmt.Println(erro)
	}

	arquivo.WriteString(site + "-online-" + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func mostrarLogs() {
	logs, _ := ioutil.ReadFile("log.txt")
	fmt.Println(string(logs))
}

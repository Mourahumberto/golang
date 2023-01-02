package main

import (
	"fmt"

	"contas/clientes"
	"contas/contas"
)

func PagarBoleto(conta verificaConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificaConta interface {
	Sacar(valor float64) string
}

func main() {
	// fmt.Println(ContaCorrente{})
	// fmt.Println(ContaCorrente{titular: "Humberto", numeroAgencia: 321, numeroConta: 123456})
	// contaDoHumberto := ContaCorrente{"Humberto", 123, 123456, 200}
	// contaDoNeto := ContaCorrente{"Neto", 123, 123456, 200}
	// fmt.Println(contaDoHumberto, contaDoNeto)

	contaDaSilvia := contas.ContaCorrente{}
	contaDoNeto := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Humberto"}}
	contaDaSilvia.Titular = clientes.Titular{Nome: "Silvia"}
	contaDaSilvia.Depositar(500)
	contaDoNeto.Depositar(900)
	fmt.Println(contaDoNeto.ObterSaldo())
	fmt.Println(contaDaSilvia.ObterSaldo())
	fmt.Println(contaDaSilvia.Sacar(100))
	fmt.Println(contaDaSilvia.ObterSaldo())

	status, valor := contaDaSilvia.Depositar(100)
	fmt.Println(status, valor)

	retorno, mensagem := contaDaSilvia.Transferir(-50, &contaDoNeto)
	fmt.Println(retorno, mensagem)
	fmt.Println(contaDaSilvia.ObterSaldo())
	fmt.Println(contaDoNeto.ObterSaldo())

	clienteHumberto := clientes.Titular{Nome: "Humberto", CPF: "012.012.012.01", Profissao: "SRE"}
	contaDoHumberto := contas.ContaCorrente{
		Titular:       clienteHumberto,
		NumeroAgencia: 123,
		NumeroConta:   123456,
	}
	fmt.Println(contaDoHumberto)

	PagarBoleto(&contaDoNeto, 30)
	fmt.Println(contaDoNeto.ObterSaldo())

}

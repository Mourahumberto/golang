package contas

import "contas/clientes"

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque efetuado com sucesso"
	} else {
		return "saldo insuficiente."
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito efetuado com sucesso", valorDoDeposito
	} else {
		return "Deposito menor que zero", valorDoDeposito
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}

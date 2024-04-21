package contas

import (
	"golang-studies-oop/titulares"
	"math/rand"
)

type ContaPoupanca struct {
	Titular               titulares.Titular
	NumeroAgencia         string
	NumeroConta, Operacao int
	saldo                 float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) (string, float64, bool) {
	if valorSaque > 0 && valorSaque < c.saldo {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso!", c.saldo, true
	} else if valorSaque < 0 {
		return "Valor de saque menor que zero!", c.saldo, false
	}

	return "Saldo insuficiente", c.saldo, false
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso", c.saldo
	}
	return "Valor do deposito menor que zero", c.saldo
}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}

func NewContaPoupanca(titular titulares.Titular, saldo float64) ContaPoupanca {
	return ContaPoupanca{
		titular,
		"0001",
		rand.Intn(899999) + 100000,
		1,
		saldo,
	}
}

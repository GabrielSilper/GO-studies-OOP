package contas

import (
	"golang-studies-oop/titulares"
	"math/rand"
)

type ContaCorrente struct {
	Titular       titulares.Titular
	NumeroAgencia string
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) (string, float64, bool) {
	if valorSaque > 0 && valorSaque < c.saldo {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso!", c.saldo, true
	} else if valorSaque < 0 {
		return "Valor de saque menor que zero!", c.saldo, false
	}

	return "Saldo insuficiente", c.saldo, false
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso", c.saldo
	}
	return "Valor do deposito menor que zero", c.saldo
}

func (c *ContaCorrente) Transferir(valorDeposito float64, contaDestino *ContaCorrente) bool {
	if _, _, realizaSaque := c.Sacar(valorDeposito); realizaSaque {
		contaDestino.Depositar(valorDeposito)
		return true
	}
	return false
}

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}

func NewContaCorrente(titular titulares.Titular, saldo float64) ContaCorrente {
	return ContaCorrente{
		titular,
		"0001",
		rand.Intn(899999) + 100000,
		saldo,
	}
}

package main

import (
	"fmt"
	"golang-studies-oop/contas"
	"golang-studies-oop/titulares"
)

type VerificarConta interface {
	Sacar(valor float64) (string, float64, bool)
}

func PagarBoleto(conta VerificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

func main() {
	contaGabriel := contas.NewContaCorrente(titulares.Titular{Nome: "Gabriel"}, 1000)
	// contaFulano := contas.NewContaCorrente(titulares.Titular{Nome: "Fulano"}, 1000)
	contaCiclano := contas.NewContaPoupanca(titulares.Titular{Nome: "Ciclano"}, 500)
	PagarBoleto(&contaGabriel, 500)
	PagarBoleto(&contaCiclano, 500)
	fmt.Println(contaGabriel)
	fmt.Println(contaCiclano)
}

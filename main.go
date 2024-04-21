package main

import (
	"fmt"
	"golang-studies-oop/contas"
	"golang-studies-oop/titulares"
)

func main() {
	contaGabriel := contas.NewContaCorrente(titulares.Titular{Nome: "Gabriel"}, 1000)
	contaFulano := contas.NewContaCorrente(titulares.Titular{Nome: "Fulano"}, 1000)

	contaGabriel.Transferir(100, &contaFulano)

	fmt.Println(contaGabriel)
	fmt.Println(contaFulano)
}

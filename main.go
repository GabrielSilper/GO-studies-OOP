package main

import "fmt"

func main() {
	type ContaCorrente struct {
		titular       string
		numeroAgencia int
		numeroConta   int
		saldo         float64
	}

	conta1 := ContaCorrente{
		titular:       "Gabriel",
		numeroAgencia: 0001,
		numeroConta:   123456,
		saldo:         100.0,
	}

	fmt.Println(conta1)
}

package main

import (
	"fmt"

	c "/contas"
)

func main() {

	contaDaSilvia := c.ContaCorrente{}
	contaDaSilvia.Titular = "Silvia"
	contaDaSilvia.Saldo = 2000

	contaDoItalo := c.ContaCorrente{}
	contaDoItalo.Titular = "Italo"
	contaDoItalo.Saldo = 5000

	genero, titular, status := contaDoItalo.Transferir(&contaDoItalo, 300, &contaDaSilvia)
	fmt.Println(genero, titular, status)
	fmt.Println("Saldo conta", contaDoItalo.Titular, " : ", contaDoItalo.Saldo)
	fmt.Println("Saldo conta", contaDaSilvia.Titular, " : ", contaDaSilvia.Saldo)

	// fmt.Println(contaDaSilvia.saldo)

	// fmt.Println(contaDaSilvia.Sacar(400))
	// fmt.Println(contaDaSilvia.saldo)

	// status, valor := contaDaSilvia.Depositar(700)
	// fmt.Println(status, valor)
	// fmt.Println(contaDaSilvia.saldo)

}

// TRABALHANDO COM STRUCT

// contaDoItalo := ContaCorrente{
// 	titular:       "Italo",
// 	numeroAgencia: 123,
// 	numeroConta:   56655,
// 	saldo:         125.50,
// }

// contaDaLuana := ContaCorrente{
// 	titular:       "Luana",
// 	numeroAgencia: 1236,
// 	numeroConta:   5605,
// 	saldo:         500.00,
// }

// fmt.Println(contaDoItalo)
// fmt.Println(contaDaLuana)

// var contaDaCris *ContaCorrente // OUTRO MODO DE PREENCHER UM STRUCT
// contaDaCris = new(ContaCorrente)
// contaDaCris.titular = "Cris"

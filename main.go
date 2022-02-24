package main

import (
	"fmt"
	"sistemabancogo/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDoDenis := contas.ContaCorrente{}
	contaDoDenis.Depositar(700)
	PagarBoleto(&contaDoDenis, 150)

	fmt.Println(contaDoDenis)

	// clienteItalo := clientes.Titular{
	// 	Nome:      "Italo",
	// 	CPF:       "123.456.789.10",
	// 	Profissao: "Desenvolvedor",
	// }
	// contaDoItalo := contas.ContaCorrente{
	// 	Titular:       clienteItalo,
	// 	NumeroAgencia: 123,
	// 	NumeroConta:   56655,
	// 	Saldo:         125.50,
	// }
	// contaDaSilvia := contas.ContaCorrente{}
	// contaDaSilvia.Saldo = 2000

	// contaDoItalo := contas.ContaCorrente{}
	// contaDoItalo.Saldo = 5000

	// genero, titular, status := contaDoItalo.Transferir(&contaDoItalo, 300, &contaDaSilvia)
	// fmt.Println(genero, titular, status)
	// fmt.Println("Saldo conta", contaDoItalo.Titular, " : ", contaDoItalo.Saldo)
	// fmt.Println("Saldo conta", contaDaSilvia.Titular, " : ", contaDaSilvia.Saldo)

	// fmt.Println(contaDaSilvia.saldo)

	// fmt.Println(contaDaSilvia.Sacar(400))
	// fmt.Println(contaDaSilvia.saldo)

	// status, valor := contaDaSilvia.Depositar(700)
	// fmt.Println(status, valor)
	// fmt.Println(contaDaSilvia.saldo)

}

// TRABALHANDO COM STRUCT

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

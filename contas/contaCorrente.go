package contas

import (
	"sistemabancogo/clientes"
)

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Transferir(contaOrigem *ContaCorrente, valorDaTransferencia float64, contaDeDestino *ContaCorrente) (string, string, string) {

	if valorDaTransferencia <= c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDeDestino.Depositar(valorDaTransferencia)
		return "Sr(a)", contaOrigem.Titular.Nome, "Transferência realizada com sucesso"
	} else {
		return "Sr(a)", contaOrigem.Titular.Nome, "Transferência não pode ser realizada por falta de saldo"
	}

}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Falha ao realizar o depósito", c.saldo
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo

}

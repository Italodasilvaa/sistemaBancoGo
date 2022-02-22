package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Transferir(contaOrigem *ContaCorrente, valorDaTransferencia float64, contaDeDestino *ContaCorrente) (string, string, string) {

	if valorDaTransferencia <= c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDeDestino.Depositar(valorDaTransferencia)
		return "Sr(a)", contaOrigem.Titular, "Transferência realizada com sucesso"
	} else {
		return "Sr(a)", contaOrigem.Titular, "Transferência não pode ser realizada por falta de Saldo"
	}

}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.Saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.Saldo
	} else {
		return "Falha ao realizar o depósito", c.Saldo
	}
}

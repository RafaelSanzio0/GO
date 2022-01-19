package contas

import "Orientacao_Objetos/banco/clientes"

//atributo que começa com minusculo = private
//atributo que começa com maisculo = public
type ContaCorrente struct {
	Titular       clientes.Titular // COMPOSIÇÃO
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

// o c* ContaCorrente é como se fosse o this ou seja quem chamou que será debitado o saldo
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podeSacar {
		c.saldo -= valorDoSaque
		return "SAQUE EFETUADO COM SUCESSO!!!"
	} else {
		return "saldo INSUFICIENTE!!!"
	}
}

// funcao com 2 retornos
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "DEPOSITO EFETUADO COM SUCESSO!!!", c.saldo
	} else {
		return "O VALOR DO DEPOSITO PRECISAR SER MAIOR QUE 0!!!", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia              // graças ao ponteiro ele realmente tira o saldo de quem chamou
		contaDestino.Depositar(valorTransferencia) // desse jeito o programa nao sabe quem é a contaDestino e por isso nao deposita o saldo (contaDestino ContaCorrente) precisa ser  (contaDestino *ContaCorrente) e colocar o & no segundo parametro na declaração do metodo
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) Obtersaldo() float64 { // como se fosse o get
	return c.saldo
}

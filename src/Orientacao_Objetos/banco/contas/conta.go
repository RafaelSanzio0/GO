package contas

type Conta interface {
	Sacar(valor float64) string // contrato (tem que seguir o mesmo parametro de
	// retorno assim como foi definido na classe contaCorrente o metodo Sacar)
}

func PagarBoleto(conta Conta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

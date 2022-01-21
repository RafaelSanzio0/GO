package contas

import "Orientacao_Objetos/banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	Saldo                                float64
}

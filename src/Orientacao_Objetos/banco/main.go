package main

import (
	"Orientacao_Objetos/banco/contas"
	"fmt"
)

func main() {

	//podemos fazer desta maneira caso queira preencher alguns campos -- JEITO MAIS CORRETO
	//clienteRafael := clientes.Titular{Nome: "Rafael", CPF: "47839318878", Profissao: "Desenvolvedor Back-end"}
	//contaRafael := contas.ContaCorrente{clienteRafael, 123, 123456, contaRafael.Depositar(123.55)}
	//fmt.Println(contaRafael.Titular, contaRafael.NumeroAgencia, contaRafael.NumeroConta, contaRafael.Obtersaldo())

	//se for preencher todos os campos fazer desta maneira
	//contaDaBruna := contas.ContaCorrente{"Bruna", 222, 111222, 200}
	//fmt.Println(contaDaBruna.Titular, contaDaBruna.NumeroAgencia, contaDaBruna.NumeroConta, contaDaBruna.Saldo)
	//fmt.Println(contaDaBruna.Sacar(201))
	//fmt.Println("Saldo atual", contaDaBruna.Saldo)

	//trabalhando com ponteiro
	//var contaDaCris *contas.ContaCorrente //informa para onde o meu new vai olhar
	//contaDaCris = new(contas.ContaCorrente)
	//contaDaCris.Titular = "Cris"
	//contaDaCris.Saldo = 500
	//fmt.Println(*contaDaCris)

	//testando interface
	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 400) // METODO DA INTERFACE

	fmt.Println(contaDaLuisa.ObterSaldo())

}

// -=-=-=-=-=-=-=-INTERFACE-=-=-=-=-=-=-=-=-
type Conta interface {
	Sacar(valor float64) string // contrato (tem que seguir o mesmo parametro de
	// retorno assim como foi definido na classe contaCorrente o metodo Sacar)
}

func PagarBoleto(conta Conta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

// -=-=-=-=-=-=-=-INTERFACE-=-=-=-=-=-=-=-=-

// declarando funcoes em GO
func SemParametro() string {
	return "Exemplo de função sem parâmetro!"
}

func UmParametro(texto string) string {
	return texto
}

func DoisParametros(texto string, numero int) (string, int) {
	return texto, numero
}

func funcaoVariadica(numeros ...int) int {
	resultadoDaSoma := 0
	for _, numero := range numeros {
		resultadoDaSoma += numero
	}
	return resultadoDaSoma
}

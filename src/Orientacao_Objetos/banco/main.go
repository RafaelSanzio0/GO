package main

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
}

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

package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	// -=-=-=-=-=01.INTRODUÇÃO AO GO E 02.TRABALHANDO COM VARIAVEIS-=-=-=-=-=
	nome := "Rafael"
	idade := 24
	versao := 1.2 // ou var versao float32 = 1.2 pois float tem duas versoes
	var variavelSemValor string

	fmt.Println("Ola", nome, "com idade", idade)
	fmt.Println("Este programa esta na versao", versao)
	fmt.Println("Nao vai imprimir nada", variavelSemValor)

	fmt.Println(reflect.TypeOf(nome)) // retonar tipo da var

	var comando int
	var comando2 int
	fmt.Scanf("%d", &comando) // insere o input do usuario na var comando
	fmt.Scan(&comando2)       // mesma coisa que o de cima so que sem o modificador

	// -=-=-=-=-=03.CONTROLE DE FLUXO-=-=-=-=-=

	if comando == 1 { // sem parenteses
		fmt.Println("1")
	} else if comando == 2 {
		fmt.Println("2")
	} else {
		fmt.Println("ERRO")
	}

	switch comando { // n precisa colocar o break
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}

	exibeIntroducao()
	comando3 := lerInput() //atribuindo funcao a variavel
	fmt.Println(comando3)

	//for { loop infinito, pois nao existe while

	//}

	// -=-=-=-=-=05.PRINCIPAIS COLECOES EM GO-=-=-=-=-=

	array := devolveEstadosDoSudeste()
	fmt.Println(array)

	exibeNomes()

	for i := 0; i < monitoramentos; i++ { // quero rodar meu codigo 5x
		for i, array := range array { // é como se fosse aquele for classico
			fmt.Println("Estou passando na posição", i,
				"do meu slice e essa posição tem o site", array)
		}
		time.Sleep(delay * time.Second)
	}
}

func exibeIntroducao() {
	fmt.Println("exemplo de funcao em go VOID")
}

func lerInput() int {
	fmt.Println("exemplo de funcao em go INTEIRO")

	var input int
	fmt.Scan(input)

	return input
}

// -=-=-=-=-=04.FAZENDO REQUISICOES PARA A WEB-=-=-=-=-=

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"
	//resp, err := http.Get(site)  funcao com mais de um retorno
	resp, _ := http.Get(site) // ignora o segundo retorno/parametro
	fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("ACESSO SUCESSO")
	} else {
		fmt.Println("ACESSO DEU ERRO")
	}

}

func devolveEstadosDoSudeste() [4]string {
	var array [4]string
	array[0] = "0"
	array[1] = "1"
	array[2] = "2"
	array[3] = "3"
	return array
}

func exibeNomes() { //SLICES
	nomes := []string{"Rafael", "Jane", "Samuel", "Jay", "Neim"}
	println(nomes)
	println("qtd intes do slice", len(nomes))
	println("capacidade", cap(nomes))

	nomes = append(nomes, "teste") // add novo elemento
	println("qtd intes do slice depois de estourar o valor", len(nomes))
	println("capacidade é dobrada apos extourar", cap(nomes))

}

func testaSite(site string) { // func com parametro sem retorno

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

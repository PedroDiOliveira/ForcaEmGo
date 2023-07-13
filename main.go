package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	jogar        bool
	respostaMenu int
)

func main() { //                                                          Main
	Introducao()
	menuDoJogo()

	jogar = true
	verificaRespostasMenu()

}

var palavraSecreta []string

func Jogando() {
	palavra := obtemNomeAleatorio()
	fmt.Println(palavra)
	for i := 0; i < len(palavra); i++ {
		palavraSecreta = append(palavraSecreta, "_")
	}
	fmt.Println(palavraSecreta)
	palavraSecreta = []string{}

}

func verificaRespostasMenu() {
	for jogar {

		switch respostaMenu {
		case 1:
			Jogando()
			fmt.Println("")
		case 2:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Resposta invalida")
		}

	}
}

func Introducao() {
	fmt.Println("******************************")
	fmt.Println("******Bem vindo a Forca*******")
	fmt.Println("******************************")
	fmt.Println("")
}
func menuDoJogo() {
	fmt.Println("(1)Novo Jogo")
	fmt.Println("(2)Fechar Programa")
	fmt.Scan(&respostaMenu)
	fmt.Println("")
}

func obtemNomeAleatorio() string {

	var arquivo = perguntaTema()

	var variavel string = "ForcaEmGo/temas/" + arquivo + ".txt"

	conteudo, err := ioutil.ReadFile(variavel)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return "erro"
	}

	// Dividir o conteúdo em linhas
	linhas := strings.Split(string(conteudo), "\n")

	// Inicializar o gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	// Obter um índice aleatório
	indiceAleatorio := rand.Intn(len(linhas))

	// Retornar o nome aleatório
	return linhas[indiceAleatorio]
}

func perguntaTema() string {

	var (
		escolhaTema int
		arquivo     string
	)

	fmt.Println("(1)Animais")
	fmt.Println("(2)CEP")
	fmt.Println("(3)Famosos")

	fmt.Scan(&escolhaTema)

	switch escolhaTema {
	case 1:
		arquivo = "animais"
	case 2:
		arquivo = "cep"
	case 3:
		arquivo = "famosos"

	default:
		fmt.Println("erro")

	}
	return arquivo
}

//TODO tentar entender o porque o numero de "_" nao esta retornando a quantidade equivalente a de letras nas palavras

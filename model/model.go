package model

import "fmt"

func Menu() (int, float64, float64) {

	var a float64 = 0
	var b float64 = 0
	var opcao int = 0

	fmt.Println("====================================================")
	fmt.Println("       CALCULADORA SIMPLES EM GO - VERSÃO 1.0       ")
	fmt.Println("====================================================")
	fmt.Println("[1] Adição")
	fmt.Println("[2] Subtração")
	fmt.Println("[3] Multiplicação")
	fmt.Println("[4] Divisão")
	fmt.Println("[5] Sair")
	fmt.Println("----------------------------------------------------")
	fmt.Println("Digite a opção desejada: ")
	fmt.Scan(&opcao)
	if opcao < 1 || opcao > 5 {
		fmt.Println("Opção inválida! Tente novamente.")
		return Menu()
	}
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&a)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&b)

	return opcao, a, b
}

func Adicao(a, b float64) float64 {
	return a + b
}

func Multiplicacao(a, b float64) float64 {
	return a * b
}

func Subtracao(a, b float64) float64 {
	return a - b
}

func Divisao(a, b float64) float64 {
	return a / b
}

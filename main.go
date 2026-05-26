package main

import (
	"fmt"

	"github.com/erik-omie/Aprendendo-Go/model"
)

func main() {

	for {

		opcao, a, b := model.Menu()

		switch opcao {
		case 1:
			resultado := model.Adicao(a, b)
			fmt.Printf("Resultado: %.2f\n", resultado)

		case 2:
			resultado := model.Subtracao(a, b)
			fmt.Printf("Resultado: %.2f\n", resultado)

		case 3:
			resultado := model.Multiplicacao(a, b)
			fmt.Printf("Resultado: %.2f\n", resultado)

		case 4:
			resultado := model.Divisao(a, b)
			fmt.Printf("Resultado: %.2f\n", resultado)

		case 5:
			println("Saindo da calculadora...")
			return

		default:
			println("Opção inválida!")
		}

	}
}

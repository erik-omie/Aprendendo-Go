package main

import "fmt"

func main() {

	var celsius float64

	for {

		fmt.Print("Informe a temperatura da água em Celsius: ") // Solicita ao usuário que insira a temperatura da água em Celsius
		fmt.Scan(&celsius)                                      // Lê a temperatura inserida pelo usuário e armazena na variável celsius

		// Condição para verificar se a temperatura é válida
		if celsius == 0 {
			fmt.Println("A temperatura da água não pode ser zero.")
			break
		}

		menssagem := avaliarTemperatura(celsius)
		fmt.Println(menssagem) // Imprime a mensagem de avaliação da temperatura da água
	}

}

func avaliarTemperatura(celsius float64) string {

	if celsius < 15 {
		return "A temperatura da água está fria."
	} else if celsius >= 15 && celsius <= 30 {
		return "A temperatura da água está agradevel!"
	} else {
		return "A temperatura da água está quente!"
	}
}

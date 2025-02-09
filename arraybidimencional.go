package main

import "fmt"

func main() {
	// Definindo o tabuleiro como uma matriz 2D
	tabuleiro := [][]string{
		[]string{"x", "-", "-"},
		[]string{"O", "X", "O"},
		[]string{"-", "O", "x"},
	}

	// Exibe a matriz de forma simples
	fmt.Println(tabuleiro)

	// Percorrendo cada linha para dinamizar ou pegar cada pedaço
	for i := 0; i < len(tabuleiro); i++ {
		// Imprime cada linha de tabuleiro em uma nova linha
		fmt.Println(tabuleiro[i])

		// Concatenar elementos da linha com espaço usando strings.Join
		// Para isso, usamos a função Join da biblioteca "strings"
		line := fmt.Sprintf("%s", tabuleiro[i]) // Formata como uma string
		fmt.Println(line)

		// Percorrendo os elementos da linha
		for j := 0; j < len(tabuleiro[i]); j++ {
			// Imprime cada elemento separado por espaço
			fmt.Print(tabuleiro[i][j], " ")
		}

		// Pula uma linha para a próxima linha
		fmt.Println()
	}
}

package main

import "fmt"

// Definindo o tipo `inteiro`
type inteiro int

// Definindo a interface `potencia`
type potencia interface {
	Quad() int
}

// Implementando o método `Quad` para o tipo `inteiro`
func (i inteiro) Quad() int {
	return int(i * i) // Retorna o quadrado do valor
}

func main() {
	// Criando uma variável do tipo `inteiro`
	var num inteiro = 3
	fmt.Println(num.Quad()) // Exibe o quadrado de `num`, que é 9

	// Criando uma variável do tipo `potencia`
	var pot potencia
	pot = num // Atribuindo um valor do tipo `inteiro` à interface `pot`
	fmt.Println(pot.Quad()) // Exibe o quadrado de `num` usando a interface, que também é 9
}

package main

import "fmt"

func main() {

	var numeros = [7]int{1, 3, 5, 6, 6, 6, 5}
	fmt.Println(numeros) // Exibe o array numeros
	//ou 
	numeros2 := [7]int{1, 3, 5, 6, 6, 6, 5}
    fmt.Println(numeros2) // Exibe o array numeros2
	//ou
	numeros3 := [...]int{1, 3, 5, 6, 6, 6, 5}
    fmt.Println(numeros3) // Exibe o array numeros3
    //ou
    numeros4 := make([]int, 7) // Cria um slice de tamanho 7
    numeros4[0] = 1
    numeros4[1] = 3
    numeros4[2] = 5
    numeros4[3] = 6
    numeros4[4] = 6
    numeros4[5] = 6
    numeros4[6] = 5
    fmt.Println(numeros4) // Exibe o slice numeros4
    //ou
    numeros5 := []int{1, 3, 5, 6, 6, 6, 5} // Cria um slice diretamente
    fmt.Println(numeros5) // Exibe o slice numeros5
    //ou
    numeros6 := make([]int, len(numeros)) // Cria um slice com o mesmo tamanho do array numeros
	// Converte o array `numeros` para slice e faz a cópia
    copy(numeros6, numeros[:])  // Usamos `numeros[:]` para criar um slice a partir do array `numeros`
    fmt.Println(numeros6) // Exibe o slice numeros6

    // Loop com slice
    for _, numero := range numeros6 {
        fmt.Println(numero) // Mostra todos os elementos do slice em cada iteração
    }

    // Loop com range e índice
    for i, numero := range numeros6 {
        fmt.Printf("Índice: %d, Valor: %d\n", i, numero)  // Mostra o índice e valor de cada elemento do slice
    }
    // ou
    for i := range numeros6 {
        fmt.Printf("Índice: %d, Valor: %d\n", i, numeros6[i])  // Mostra o índice e valor de cada elemento do slice
    }
    // ou
	for numero := range numeros6 {
        fmt.Println(numero) // Mostra todos os elementos do slice em cada iteração
    }
    // ou
    for i, numero := range numeros6 {
        fmt.Printf("Índice: %d, Valor: %d\n", i, numero) // Mostra o índice e valor de cada elemento do slice
    }
    // ou
    for i := range numeros6 {
        fmt.Printf("Índice: %d, Valor: %d\n", i, numeros6[i])  // Mostra o índice e valor de cada elemento do slice
    }
    // ou
    numeros7 := []int{1, 3, 5, 6, 6, 6, 5} // Cria um novo slice
    numeros7 = append(numeros7, 7) // Adiciona um novo valor ao final do slice
    fmt.Println(numeros7)  // Exibe o slice com o novo número adicionado
}

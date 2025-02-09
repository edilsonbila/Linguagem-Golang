package main
import "fmt"
func main() {

	var numeros = [7]int {1,3,5,6,6,6,5}
	//slice pega uma pare do codigo eu quis pegar os 3 no meio e quando digitate o ultimo vavolr faz se o indice +1

	fmt.Println(numeros[2:5])// IMPRIME   5,6,6
//verificar tamanho
fmt.Println("len=%d    cap=%d",len(numeros), cap(numeros)) //len = 3  cap=7
}
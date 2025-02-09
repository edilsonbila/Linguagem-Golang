package main
import "fmt"


func soma(x int, y int) int {
	return x + y} 

//funcao que retorna mais de um valor por isso declaramos 2 tipos de retorno
func calcular(a int) (int ,int){
	var quadrado int = a * a
	var cubo int = a * a * a
	return quadrado, cubo
	
}


/**valores de retorno renomeado
func valores (a int) (quadrado int, cubo int) {

   quadrado = a * a
   cubo = a * a * a
   e nao precisa escrever  apalavra reuturn assim funciona
}


**/

func main() {
	fmt.Println(soma(1, 2))
	fmt.Println(calcular(3)) //retorna 9 e 27
}
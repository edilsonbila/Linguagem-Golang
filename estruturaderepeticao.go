package main
import "fmt"

func main() {

	// for que soma os numeros de 1 a dez
	var soma int = 0
	for i := 1; i <= 10; i++ {
        soma += i
    }

 fmt.Println("A soma dos numeros de 1 a 10 e: ", soma)// 55 

/**
soma dos multiplos de 2 usando for  como while
 var soma int = 2; 
 for ; soma <600;  {
  soma += soma
}
  fmt.Println("A soma dos multiplos de 2 e: ", soma )
**/








}
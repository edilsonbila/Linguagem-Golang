package main
import "fmt"


func main (){

var numero = [] int {1,2,3,4,5}

for i := 0; i < len(numero);i++{
	fmt.Println(numero)
}

//usando range
for i := range numero{

	fmt.Println(numero[i])
}





}
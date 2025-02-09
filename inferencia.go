package main
import "fmt"

func main() {
var i int64
var j = i //nos nao definimos o tipo de j logo j ira herdar o tipo de i
fmt.Printf("TIPO: %T VALOR: %V\n", j, j)
//logo ira imprimir que o tipo e int o valor omo nao foi declarado ele entende como zero

/**mas se

var i = 35 o numero e inteiro logo sabemos que o tipo e inteiro e valor e 35
var i = 23.3 + 43i
ao impremir para saber o tipo de dado ele dira que e um tipo complex  de valor 23.3 + 43i
**/



}
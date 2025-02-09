package main
import "fmt"


func main (){

 var a int = 24

 if a > 10 {
	fmt.Println("a variavel a e maior que 10")
 }else if a > 5 {

    fmt.Println("a variavel a esta entre 6 e dez")
 }else{
	  fmt.Println("a variavel a nao e maior que 10 nem 5")
 }
// no go o else fica logo na linha de feichamento do if
/** no if podemos definir variaveis

if a :=11; a > 10{
    fmt.Println("a variavel a e maior que 10")}
 mas a variavel so pode ser usada na estrutur de condicao fora a variavel nao estara definid**/
 

}
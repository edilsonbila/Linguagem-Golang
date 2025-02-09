package main
import "fmt"


func main(){
	//as variaveis podem ser definidas da seguinte forma var idade =23 e var nome String="meuNome".
	var idade int =23
	var nome ="meuNome"
	//o "p sempre deve ser maiusculo para a impressao ".
	fmt.Println("o meu nome e ",nome," e tenho ",idade)

}
/**mas quando temos valores ou variaveis do mesmo tipo podedos declarar da seguinte forma
var idade ,numero = 23,45 
**/
/**mas existe uma forma de declarar diferente tambem que e 
var(
   numero = 23
   nome = "meuNome"


)



ou Ainda

numero :=23
nome:= "meuNome"




ou
numero,nome := 23,"meuNome"


**/
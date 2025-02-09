package main
import "fmt"


func main (){
	//a chave e string e vai receber ints
 var notas map [string]int
 notas = make(map[string]int)

 notas["ana"]=8

 notas["antonio"]=15

 fmt.Println(notas) //mostrs ana 8 e antonio 15
 fmt.Println(notas["ana"]) //retorna 8

 valor,existe :=notas["julio"]
  fmt.Println(valor)//retorna zero porque nachave julio naotem valor associado
 fmt.Println(existe)//se achave tiver recebendo um valor retorna true caso contrario false


}

/**outra forma


var notas map[string]int
notas =map[atring]int{
  "ana":9,
  "maria":10,
  }

  fmt.println(notas)




**/



/**stict com map

type perfil struct{

altura float64
ativo bool
profissao  string

}


func main(){
var perfis map [string]perfil = make(map[string]perfil)

perfis["joao"] = perfil{
1.74,true,"medico",
}

fmt.Println(perfis )
}







**/
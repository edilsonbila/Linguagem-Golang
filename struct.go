package main


import "fmt"

func main (){

	type posicao struct{

		x int
		y int
	}
   fmt.Println(posicao{32,23})// valor de x=32	,y=23


   //MODIFICAR STRICT
   pos := posicao{32,23}
   pos.x = 45
   fmt.Println(pos.x)// o x passa a ser 45
}
	
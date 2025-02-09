package main
import "fmt"

func main (){
  a := 32
  p := &a //referencia
  fmt.Println(p)// ira imprimir valor decimal
  
     *p = 42
     fmt.Println(a) // ira imprimir 42
     /**
     // ou
     *p += 1
     fmt.Println(a) // ira imprimir 43
     // ou
     *p = *p * 2
     fmt.Println(a) // ira imprimir 84


	 // se imprimir *p
	 	fmt.Println(*p)
   */



  



}
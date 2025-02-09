package main
import "fmt"


func main() {

var nome string ="joao"

 switch nome  {
	case "aNA":
       fmt.Println("E ANA")
	 case "JOAO" :
	   fmt.Println("E JOAO") 
	 default: 
	     fmt.Println("Nome desconhecido")

	
 }

  fmt.Println(nome)

  /**switch com reacao de if "switch true"
          
         var nota int = 8
      
           switch true {

case nota > 9 :
	 fmt.Println("OTIMO")
case nota > 7 :
	 fmt.Println("muitoBOM")
case  nota > 6 :
	 fmt.Println("BOM")
	 default:
		fmt.Println("RUIM")
		}	 	 


  **/


}
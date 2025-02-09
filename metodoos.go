package main
import "fmt"

type pessoa struct {

	nome,sobrenome string
}



func main(){
	alguem := pessoa {"jose","rosa"}
	fmt.Println(alguem)


}
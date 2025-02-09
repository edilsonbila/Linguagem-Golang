package main
import "fmt"

func main() {
defer fmt.Println("oi");// se colocarmos o defer ele ira priorizar a proxima execucao
fmt.Println("tudo bem?");

}
//ira imprimir primeiro o tudo bem?  depois o oi

package main 
import "fmt"
 
func main() {
var aberto = true
var altura = 1.34
fmt.Printf("TIPO: %T VALOR: %v",altura, altura)
// AO IMPRIMIR DEVE MOSTRAR O TIPO E O VALOR 

fmt.Printf("TIPO: %T VALOR: %v",aberto,aberto)
//tambem mostra tipo bool valor true

}

/**
existem os seguintes tipos de dados para go
int int8 int16 int32 int64
uint tambem e inteiro mas so inteiro positivo uint8 uint16 uint32 uint
bool que e boleano

float32 float64
rune
complex64 complex128
**/
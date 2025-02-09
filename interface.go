package main
import "fmt"


type Geometrica interface{

	area () float64

}

type Quadrado struct {
	lado float64

	
}


type circulo struct {
	raio float64

	
}


func(q Quadrado) area() float64{
	return q.lado * q.lado
}
func main(){
	
var  g Geometrica 
g= 	Quadrado{3}
fmt.Printf(" a area do quadrado e %v ", g.area())
}
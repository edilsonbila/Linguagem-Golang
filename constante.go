package main
import "fmt"
import"math"

func main() {
	const pi = 3.14// todas constates nao podem ser redefinidas
   const numero = math.Max(3,5)//nao se define constande a esse tipo de valor  que ainda pode ser processado
   var numero = math.Max(3,5)//essa sim mas nao e constante	


fmt.Println(pi)
}
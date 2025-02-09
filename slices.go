package main
import "fmt"

func main() {
	var NOMES [] string;

	var NOMES2= append(NOMES, "jimmy")
	NOMES2 = append(NOMES2, "mary")
	fmt.Println(NOMES)
	fmt.Println("len= %d  cap= %d",len(NOMES2),cap(NOMES2))
}
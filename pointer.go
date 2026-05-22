package main
import(
	"fmt"
)

var number int = 45
func main(){

	anotherNumber := &number

	fmt.Printf("anothernumber's value is %d", *anotherNumber)
}
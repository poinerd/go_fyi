package main

import(
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0{
		return 0, errors.New("cannot divide by 0")
	}

	return a / b, nil
}

func celebrateBirthday(a *int){
	*a = *a + 1
	fmt.Printf("The new age of this person is %d", *a)
}

func main(){

	age := 25

	result, err := divide(10,0)

	if err!= nil {
		fmt.Println("Error occured", err)
	}
     
	fmt.Println("Result" , result)
	celebrateBirthday(&age)

}


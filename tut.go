package main

import (
	"fmt"
)

// This is the explicit way of declaring stuff in go
var isUserOnline bool = true

func swapValues(a *int, b *int) {
	p := *a  // 1. Save the value of 'a' (1) into 'p'
	*a = *b  // 2. Overwrite 'a' with the value of 'b' (2)
	*b = p   // 3. Overwrite 'b' with the saved value of 'p' (1)

	fmt.Printf("Inside swapValues: a is now %d and b is %d\n", *a, *b)
}

type Car struct{
	available bool
}

func (d Car) getAvailability () bool{
	return d.available
}

func main() {
	num1 := 1
	num2 := 2

	swapValues(&num1, &num2)

	myCar := Car{ available: true}
	carStatus := myCar.getAvailability()
	
	fmt.Printf("After swap in main: num1 = %d, num2 = %d\n", num1, num2)
	fmt.Printf("The car's availability is %t\n", carStatus)
}
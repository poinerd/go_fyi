// Slices, Arrays and Maps
package main

import (
	"fmt"
	"runtimem, ")

func loopOverArray(arr []int){
	for i := 0; i<len(arr) ; i++{
		fmt.Println(arr[i])
	}
}

func main(){
	cores := runtime.NumCPU()
	fmt.Printf("Your computer has %d physical CPU cores.\n", cores)

	// anotherArray := [3]int{78,99,20}
	// aSlice := make([]int, 0, 10)

	// class25 := make(map[string]int)
	// class25["number"] := 34
	myArray := [5]int{1, 2, 4, 5, 9} // Numbers must be in double quotations
	loopOverArray(myArray[:])
	// This "=" is used to updating an already created variable

	threads := runtime.GOMAXPROCS(0) 
	fmt.Printf("Go has allocated a pool of %d OS threads for your code.\n", threads)
}
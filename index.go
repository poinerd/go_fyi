package main

import "fmt"

 /// Go doesnt have class, extends. You build complex stuff by combining small stuffs
 
 // Go is fast, complies directly ti binary.

 // Go ships a single static binary with an embedded garbage collector

 // What is an embedded garbage collecor?

 // Go is single threaded

//  var counter int = 0


func add (a,b int)(int){
	return a+b
}

func main(){
      counter := 0  // Go infers this to be an integer

	//   Use := when you want to create a variable for the first time inside a function without typing out var. Use = when you just want to update a variable that already exists.

	for i := 0; i<5; i++ {
		if i == 3 {
			fmt.Printf("Hello there\n")
		} else{
			fmt.Printf("Whats up?\n")
		}
	}


	for counter < 10 {
		fmt.Printf("counter is less than 10, counter is %d \n", counter)
		counter ++
	}
    
	result := add(2,3)
	fmt.Printf("%d", result)
	


	// Go has pointers but no pointer arithmetic (you cant do ptr++ It used pointers
	// purely for memeoru efficiency and mutability

}
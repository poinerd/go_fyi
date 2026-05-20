package main

import (
	"fmt"
)

func runLoop (a int){
	for i:= 0; i<=a ; i++{
		fmt.Printf("%d\n", i)
	}

}

func runLoopWithWhile (a int){
	counter := 0

	for counter < a {
		fmt.Printf("%d\n", a)
		counter ++
	}
}

func main(){
	runLoopWithWhile(20)
   
}
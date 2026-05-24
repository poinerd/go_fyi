package main

import (
	"fmt"
)



func main(){
  // This must be inside of a function, it cant be global
	d := 300
    p := &d // p is now a pointer to d because it stores the address of d
	// *p = 100

	fmt.Println(p)

}

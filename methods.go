package main

import(
	"fmt"
)
type animal struct{
	name string
	legs int
	sound string
}

func newAnimal (n string, l int, s string) animal{

	return animal{
	name: n,
	legs : l,
	sound : s,}

}

func (a animal) getName () string{
	name := a.name
	return name
}	

func (a animal) getLegs() int{
	legs := a.legs
	return legs
}

func (a animal) getSound () string{
	 sound := a.sound
	return sound
}
func main(){

	// myAnimal := animal{name : "Bird", legs : 2, sound : "meow"}
	newStuff := newAnimal("Dragon", 2, "Roar")
	fmt.Println(newStuff)
	fmt.Println(newStuff.getName())
	fmt.Println(newStuff.getLegs())
	fmt.Println(newStuff.getSound())

}
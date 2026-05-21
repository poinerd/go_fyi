package main

import (
	"fmt"
	"time"
)

type birds struct{
	name string
	sound  string
	legs int
}

func (bird birds) getBirdName() string{
	return bird.name
}

func (bird birds) getBirdLegs() int{
	return bird.legs
}

func (bird birds) getBirdSound() string{
	return bird.sound
}

func main(){

	myBird := birds{name: "Owl", sound: "Ohh", legs: 2}
	birdName := myBird.getBirdName()
	birdLegs := myBird.getBirdLegs()
	birdSound := myBird.getBirdSound()

	fmt.Printf("The bird's name is %s, it smakes the sound %s and it has %d legs", birdName, birdSound, birdLegs)
}
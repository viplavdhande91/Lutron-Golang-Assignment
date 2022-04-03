package main

import (
	"fmt"
)

//INTERFACE
type Planet interface {
	Name() string
	Mass() int
}

//STRUCTURE DEFINITION
type planetDetails struct {
	name string
	mass int
}

func (o planetDetails) Name() string {

	return o.name
}

func (o planetDetails) Mass() int {

	return o.mass
}

func planetNameDisplay(obj Planet) {

	fmt.Println("Name: ", obj.Name())
	fmt.Println("Mass: ", obj.Mass())

}

func main() {
	p1 := planetDetails{
		name: "pluto",
		mass: 5000,
	}
	p2 := planetDetails{
		name: "earth",
		mass: 6000,
	}
	planetNameDisplay(p1)
	planetNameDisplay(p2)

}

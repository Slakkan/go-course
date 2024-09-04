package main

import "fmt"

type Animal interface {
	Says() string
	numberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) numberOfLegs() int {
	return 4
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func (g *Gorilla) Says() string {
	return "Ugh"
}

func (g *Gorilla) numberOfLegs() int {
	return 2
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepard",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name: "Jock",
		Color: "grey",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.numberOfLegs(), "legs")
}
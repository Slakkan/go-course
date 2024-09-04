package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	var whatToSay string 
	var i int

	whatToSay = "Potatoes"
	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	whatWasSaid, otherThingSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, otherThingSaid)
}

func saySomething() (string, string) {
	return "somthing", "else"
}
package main

import "log"

func main() {
	var myString string
	myString = "green"

	log.Println("myString is set to", myString)

	changeUsingPointer(&myString)

	log.Println("After calling changeUsingPointer myString is set to", myString)
}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
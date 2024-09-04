package main

import (
	"log"
)

// STRUCTS

type myStruct struct  {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}

// // Capital letters means the variable is public, therefore you can use it in other packages
// type User struct {
// 	FirstName string
// 	LastName string
// 	PhoneNumber string
// 	Age int
// 	BirthDate time.Time
// }

// func main () {
// 	user := User{
// 		FirstName: "Lisandro",
// 		LastName: "Pastinante",
// 	}

// 	log.Println(user.FirstName, user.LastName, "BirthDate:", user.BirthDate)
// }


// SHADOWING

// var s = "seven"

// func main() {
// 	s := "eight"

// 	log.Println("s in main is", s)

// 	saySomething("xxx")
// }

// func saySomething(s string) (string, string) {
// 	log.Println("s from the saySomething func is", s)
// 	return s, "World"
// }
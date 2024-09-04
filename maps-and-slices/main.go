package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName string
}

func main() {
	var mySlice []string
	mySlice = append(mySlice, "Lisandro")
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Mary")

	sort.Strings(mySlice)

	log.Println(mySlice)

	numbers := []int{1,2,3,4,5,6,7,8,9,10}

	log.Println(numbers)

	log.Println(numbers[0:2])
}

// MAPS
// func main() {
// 	myMap := make(map[string]User)
	
// 	me := User {
// 		FirstName: "Lisandro",
// 		LastName: "Pastinante",
// 	}

// 	myMap["me"] = me
	
// 	log.Println(myMap["me"].FirstName, myMap["me"].LastName)
// }


// func main() {
// myMap := make(map[string]string)

// myMap["dog"] = "Samson"

// log.Println((myMap["dog"]))
// }
package main

import (
	"log"

	"github.com/slakkan/my-package/helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <- intChan
	log.Println(num)
}

// import (
// 	"log"
// 	"github.com/slakkan/my-package/helpers"
// )

// func main() {
// 	var myVar helpers.SomeType
// 	myVar.TypeName = "Some name"
// 	log.Println(myVar.TypeName)
// }

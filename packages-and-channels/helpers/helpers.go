package helpers

import (
	"math/rand"
)

func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}

// type SomeType struct {
// 	TypeName   string
// 	TypeNumber int
// }
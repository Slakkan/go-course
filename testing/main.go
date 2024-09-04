package main

import (
	"errors"
	"log"
)

// coverage command, run on git-bash: go test -coverprofile=coverage.out && go tool cover -html=coverage.out

func main() {
	result, err := divide(100.0, 0)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Result:", result)
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("cannot divide by 0")
	}

	result = x / y
	return result, nil
}
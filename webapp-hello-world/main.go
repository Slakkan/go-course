package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = 3000

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home\n")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About\n")

	sum := addValues(2,2)
	fmt.Fprintf(w, "And 2+2 is %d", sum)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	fmt.Fprintf(w,"%f divided by %f is %f", 100.0, 0.0, f)
}

// addValues adds two integers and returns the sum
func addValues(x, y int) int {
	return x + y
}

func divideValues (x, y float32) (float32, error){
	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x/y
	return result, nil
}

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	
	fmt.Printf("Listening requests from http://localhost:%d\n", portNumber)
	http.ListenAndServe(fmt.Sprintf(":%d", portNumber),nil)
}
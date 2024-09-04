package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		n, err := fmt.Fprintf(w, "Hello, world!")
		fmt.Println(fmt.Sprintf("Bytes written: %d", n))
		if err != nil {
			fmt.Println(err)
		}
	})

	http.ListenAndServe(":3000",nil)
}
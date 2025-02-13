package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting the server at port: 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic("Error while running the server")
	}
}

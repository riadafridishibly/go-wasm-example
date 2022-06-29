package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening to: http://localhost:8080/")
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("assets")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}

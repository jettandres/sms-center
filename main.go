package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := NewSmsServer()
	fmt.Println("server running...")
	fmt.Println(http.ListenAndServe(":5000", server))
}

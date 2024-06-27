package main

import (
	"fmt"
	"net/http"
)

func main() {
	store := NewInMemoryStore()
	server := NewSmsServer(store)
	port := ":5000"
	fmt.Printf("server running at port localhost%s", port)
	fmt.Println(http.ListenAndServe(port, server))
}

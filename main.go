package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(SmsServer)
	fmt.Println("server running...")
	fmt.Println(http.ListenAndServe(":5000", handler))
}

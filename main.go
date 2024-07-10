package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
)

func main() {
	store := NewSqliteStore()
	server := NewSmsServer(store)
	auth := NewAuth(server)

	port := ":5000"
	fmt.Printf("server running at port localhost%s\n", port)
	fmt.Println(http.ListenAndServe(port, auth))
}

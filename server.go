package main

import (
	"net/http"
)

func SmsServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

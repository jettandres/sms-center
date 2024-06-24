package main

import (
	"net/http"
)

type Sms struct {
	Id          string `json:"id"`
	Inserted_at string `json:"inserted_at"`
	From        string `json:"from"`
	To          string `json:"to"`
	Body        string `json:"body"`
}

type SmsServer struct {
	http.Handler
}

func NewSmsServer() *SmsServer {
	return new(SmsServer)
}

func (s *SmsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()
	router.Handle("/sms", http.HandlerFunc(handleSms))
	router.ServeHTTP(w, r)
	w.WriteHeader(http.StatusNotFound)
}

func handleSms(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

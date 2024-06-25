package main

import (
	"encoding/json"
	"net/http"
)

type Sms struct {
	Id          string `json:"id"`
	Inserted_at string `json:"inserted_at"`
	From        string `json:"from"`
	To          string `json:"to"`
	Body        string `json:"body"`
}

type ResponseData struct {
	Sms []Sms `json:"sms"`
}

type ServerResponse struct {
	Status string       `json:"status"`
	Data   ResponseData `json:"data"`
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
}

func handleSms(w http.ResponseWriter, r *http.Request) {
	resp := ServerResponse{
		Status: "success",
		Data: ResponseData{
			Sms: []Sms{},
		},
	}

	body, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

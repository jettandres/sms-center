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

type AllSmsData struct {
	Sms []Sms `json:"sms"`
}

type GetAllSmsResponse struct {
	Status string     `json:"status"`
	Data   AllSmsData `json:"data"`
}

type SmsServer struct {
	store Store
	http.Handler
}

func NewSmsServer(store Store) *SmsServer {
	server := new(SmsServer)
	server.store = store
	return server
}

func (s *SmsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()

	router.HandleFunc("GET /sms", s.handleGetAllSms)
	router.HandleFunc("GET /sms/{mobileNumber}", s.handleGetSms)

	router.ServeHTTP(w, r)
}

func (s *SmsServer) handleGetAllSms(w http.ResponseWriter, r *http.Request) {
	resp := GetAllSmsResponse{
		Status: "success",
		Data: AllSmsData{
			Sms: s.store.GetAllSms(),
		},
	}

	body, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

func (s *SmsServer) handleGetSms(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

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

type SmsPayload struct {
	From string `json:"from"`
	To   string `json:"to"`
	Body string `json:"body"`
}

type AllSmsData struct {
	Sms []Sms `json:"sms"`
}

type SmsData struct {
	Sms Sms `json:"sms"`
}

type GetAllSmsResponse struct {
	Status string     `json:"status"`
	Data   AllSmsData `json:"data"`
}

type GetAllSmsFromNumberResponse struct {
	Status string     `json:"status"`
	Data   AllSmsData `json:"data"`
}

type GetSmsFromNumberResponse struct {
	Status string  `json:"status"`
	Data   SmsData `json:"data"`
}

type PostSmsResponse struct {
	Status string  `json:"status"`
	Data   SmsData `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
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
	router.HandleFunc("GET /sms/{mobileNumber}", s.handleGetAllSmsFromNumber)
	router.HandleFunc("GET /sms/{mobileNumber}/{id}", s.handleGetSmsFromMobileNumber)

	router.ServeHTTP(w, r)
}

func (s *SmsServer) handleGetAllSms(w http.ResponseWriter, r *http.Request) {
	resp := GetAllSmsResponse{
		Status: "success",
		Data: AllSmsData{
			Sms: s.store.GetAllSms(),
		},
	}

	body, err := json.Marshal(resp)
	if err != nil {
		handleError(err, http.StatusInternalServerError, w)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

func (s *SmsServer) handleGetAllSmsFromNumber(w http.ResponseWriter, r *http.Request) {
	mobileNumber := r.PathValue("mobileNumber")
	resp := GetAllSmsFromNumberResponse{
		Status: "success",
		Data: AllSmsData{
			Sms: s.store.GetAllSmsFromNumber(mobileNumber),
		},
	}

	body, err := json.Marshal(resp)
	if err != nil {
		handleError(err, http.StatusInternalServerError, w)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *SmsServer) handleGetSmsFromMobileNumber(w http.ResponseWriter, r *http.Request) {
	mobileNumber := r.PathValue("mobileNumber")
	resp := GetSmsFromNumberResponse{
		Status: "success",
		Data: SmsData{
			s.store.GetSmsFromNumber(mobileNumber),
		},
	}

	body, err := json.Marshal(resp)
	if err != nil {
		handleError(err, http.StatusInternalServerError, w)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func handleError(err error, statusCode int, w http.ResponseWriter) {
	resp := ErrorResponse{
		Status:  "error",
		Message: err.Error(),
	}

	body, _ := json.Marshal(resp)

	w.WriteHeader(statusCode)
	w.Write(body)
}

package main

import (
	"encoding/json"
	"net/http"
)

type Sms struct {
	Id          string `json:"id"`
	Inserted_at string `json:"inserted_at"`
	Sender      string `json:"sender"`
	Receiver    string `json:"receiver"`
	Body        string `json:"body"`
}

type SmsPayload struct {
	Receiver string `json:"receiver"`
	Body     string `json:"body"`
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

type InsertSmsResponse struct {
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
	router.HandleFunc("GET /sms/{sender}", s.handleGetAllSmsFromSender)
	router.HandleFunc("GET /sms/{sender}/{id}", s.handleGetSmsFromSender)
	router.HandleFunc("POST /sms/{sender}", s.handleInsertSms)

	router.ServeHTTP(w, r)
}

func (s *SmsServer) handleGetAllSms(w http.ResponseWriter, r *http.Request) {
	data, err := s.store.GetAllSms()
	if err != nil {
		handleError(err, http.StatusBadRequest, w)
	}

	resp := GetAllSmsResponse{
		Status: "success",
		Data: AllSmsData{
			Sms: data,
		},
	}

	body, err := json.Marshal(resp)
	if err != nil {
		handleError(err, http.StatusInternalServerError, w)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

func (s *SmsServer) handleGetAllSmsFromSender(w http.ResponseWriter, r *http.Request) {
	sender := r.PathValue("sender")

	data, err := s.store.GetAllSmsFromSender(sender)
	if err != nil {
		handleError(err, http.StatusBadRequest, w)
	}

	resp := GetAllSmsFromNumberResponse{
		Status: "success",
		Data: AllSmsData{
			Sms: data,
		},
	}

	body, err := json.Marshal(resp)
	if err != nil {
		handleError(err, http.StatusInternalServerError, w)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *SmsServer) handleGetSmsFromSender(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	data, err := s.store.GetSmsById(id)
	if err != nil {
		handleError(err, http.StatusBadRequest, w)
	}

	resp := GetSmsFromNumberResponse{
		Status: "success",
		Data: SmsData{
			data,
		},
	}

	body, err := json.Marshal(resp)
	if err != nil {
		handleError(err, http.StatusInternalServerError, w)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *SmsServer) handleInsertSms(w http.ResponseWriter, r *http.Request) {
	sender := r.PathValue("sender")

	var payload SmsPayload
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		handleError(err, http.StatusInternalServerError, w)
	}

	data, err := s.store.InsertSms(payload.Receiver, sender, payload.Body)
	if err != nil {
		handleError(err, http.StatusBadRequest, w)
	}

	resp := InsertSmsResponse{
		Status: "success",
		Data: SmsData{
			data,
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

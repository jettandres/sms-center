package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Auth struct {
	http.Handler
}

type AuthResponse struct {
	Status string `json:"status"`
}

func NewAuth(handlerToWrap http.Handler) *Auth {
	return &Auth{handlerToWrap}
}

func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authUrl := os.Getenv("AUTH_VALIDATE_URL")
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, authUrl, nil)
	req.Header = r.Header

	if err != nil {
		panic(err.Error())
	}

	res, err := client.Do(req)

	if err != nil {
		handleError(err, http.StatusUnauthorized, w)
	}

	var authResp AuthResponse
	err = json.NewDecoder(res.Body).Decode(&authResp)

	if err != nil || authResp.Status != "success" {
		handleError(err, http.StatusUnauthorized, w)
	}

	a.Handler.ServeHTTP(w, r)
}

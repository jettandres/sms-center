package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("200 OK /sms", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/sms", nil)
		response := httptest.NewRecorder()

		store := NewInMemoryStore()

		server := NewSmsServer(store)
		server.ServeHTTP(response, request)

		assertStatusOk(t, response)

		var body GetAllSmsResponse
		err := json.NewDecoder(response.Body).Decode(&body)

		if err != nil {
			t.Errorf("unable to parse response from server")
		}
	})

	t.Run("200 OK /sms/:mobile-number", func(t *testing.T) {
		mobileNumber := "0916123456"
		request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/sms/%s", mobileNumber), nil)
		response := httptest.NewRecorder()

		store := NewInMemoryStore()

		server := NewSmsServer(store)
		server.ServeHTTP(response, request)

		assertStatusOk(t, response)
	})
}

func assertStatusOk(t *testing.T, response *httptest.ResponseRecorder) {
	t.Helper()
	if response.Result().StatusCode != http.StatusOK {
		t.Errorf("incorrect status code, want %d, got %s", http.StatusOK, response.Result().Status)
	}

}

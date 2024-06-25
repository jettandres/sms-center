package main

import (
	"encoding/json"
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

		if response.Result().StatusCode != http.StatusOK {
			t.Errorf("incorrect status code, want %d, got %s", http.StatusOK, response.Result().Status)
		}

		var body GetAllSmsResponse
		err := json.NewDecoder(response.Body).Decode(&body)

		if err != nil {
			t.Errorf("unable to parse response from server")
		}
	})
}

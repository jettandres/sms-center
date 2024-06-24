package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("200 OK /sms", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/sms", nil)
		response := httptest.NewRecorder()

		SmsServer(response, request)

		if response.Result().StatusCode != http.StatusOK {
			t.Errorf("incorrect status code, want %d, got %s", http.StatusOK, response.Result().Status)
		}
	})
}

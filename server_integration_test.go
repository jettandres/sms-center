package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	store := NewInMemoryStore()
	server := NewSmsServer(store)

	fromNumber := "0906765432"
	toNumber := "0916123456"

	server.ServeHTTP(httptest.NewRecorder(), newInsertSmsRequest(fromNumber, toNumber, "hello integ"))
	server.ServeHTTP(httptest.NewRecorder(), newInsertSmsRequest(fromNumber, toNumber, "hello again integ"))
	server.ServeHTTP(httptest.NewRecorder(), newInsertSmsRequest(fromNumber, toNumber, "hello once more integ"))

	t.Run("retrieve ALL stored messages", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/sms", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var apiResp GetAllSmsResponse
		err := json.NewDecoder(response.Body).Decode(&apiResp)

		if err != nil {
			t.Errorf("Unable to parse response. Error: %s", err.Error())
		}

		if len(apiResp.Data.Sms) != 3 {
			t.Error("Unable to retrieve any messages. Must have the 3 messages inserted")
		}
	})

	t.Run("retrieve ALL messages of a number", func(t *testing.T) {
		fromAnotherNumber := "0909696969"
		server.ServeHTTP(httptest.NewRecorder(), newInsertSmsRequest(fromAnotherNumber, toNumber, "hello this is Yeji"))

		request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/sms/%s", fromAnotherNumber), nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var apiResp GetAllSmsFromNumberResponse
		err := json.NewDecoder(response.Body).Decode(&apiResp)

		if err != nil {
			t.Errorf("Unable to parse response. Error: %s", err.Error())
		}

		if len(apiResp.Data.Sms) != 1 {
			t.Errorf("Expecting just 1 message from number %s, got %d message/s", fromAnotherNumber, len(apiResp.Data.Sms))
		}
	})

	t.Run("view a specific message", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newInsertSmsRequest(fromNumber, toNumber, "this is the final message"))

		var insertApiResp InsertSmsResponse
		err := json.NewDecoder(response.Body).Decode(&insertApiResp)

		if err != nil {
			t.Errorf("Unable to parse response. Error: %s", err.Error())
		}

		id := insertApiResp.Data.Sms.Id
		request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/sms/%s/%s", fromNumber, id), nil)

		getResponse := httptest.NewRecorder()
		server.ServeHTTP(getResponse, request)

		var getSmsFromNumResp GetSmsFromNumberResponse
		err = json.NewDecoder(getResponse.Body).Decode(&getSmsFromNumResp)

		if err != nil {
			t.Errorf("Unable to parse response. Error: %s", err.Error())
		}

		if getSmsFromNumResp.Data.Sms.Id != id {
			t.Errorf("Expecting to retrieve SMS with id (\"%s\"), got SMS with id (\"%s\")", id, getSmsFromNumResp.Data.Sms.Id)
		}
	})
}

func newInsertSmsRequest(fromNumber string, toNumber string, body string) *http.Request {
	reqBody := SmsPayload{
		From: fromNumber,
		Body: body,
	}
	payload, _ := json.Marshal(reqBody)
	return httptest.NewRequest(http.MethodPost, fmt.Sprintf("/sms/%s", toNumber), strings.NewReader(string(payload)))
}

package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"summar/server/stores"
	"summar/server/utils"
	"testing"
)

func TestHandlers(t *testing.T) {
	mockStore := stores.NewMockStore()

	handlers := Handlers{
		Store: mockStore,
	}

	t.Run("Test signup handler", func(t *testing.T) {
		reqBody := utils.JSONMarshal(&SignupRequest{
			Email:    "test@test.com",
			Password: "123",
		})

		req, err := http.NewRequest("POST", "/signup", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler := ToHttpHandlerFunc(handlers.SignupHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler errored out: got %v want %v", status, http.StatusOK)
		}

		var resBody SignupResponse
		if err := json.Unmarshal(rr.Body.Bytes(), &resBody); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test login handler", func(t *testing.T) {
		reqBody := utils.JSONMarshal(&LoginRequest{
			Email:    "test@test.com",
			Password: "123",
		})

		req, err := http.NewRequest("POST", "/login", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler := ToHttpHandlerFunc(handlers.LoginHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler errored out: got %v want %v", status, http.StatusOK)
		}

		var resBody HandlerResponse
		if err := json.Unmarshal(rr.Body.Bytes(), &resBody); err != nil {
			t.Fatal(err)
		}
	})
}

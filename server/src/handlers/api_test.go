package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"summar/server/stores"
	"summar/server/utils"
	"testing"

	"github.com/google/uuid"
)

func TestHandlers(t *testing.T) {
	mockStore := stores.NewMockStore()

	handlers := Handlers{
		Store: mockStore,
	}

	var userId uuid.UUID

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

		t.Logf("signup response:\n%+v\n", resBody)

		userId = resBody.User.Id
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

		t.Logf("login response:\n%+v\n", resBody)
	})

	t.Run("Test get user handler", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/login", nil)
		if err != nil {
			t.Fatal(err)
		}

		ctx := context.WithValue(req.Context(), "userId", userId)
		req = req.WithContext(ctx)

		rr := httptest.NewRecorder()

		handler := ToHttpHandlerFunc(handlers.GetUserHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler errored out: got %v want %v", status, http.StatusOK)
		}

		var resBody GetUserResponse
		if err := json.Unmarshal(rr.Body.Bytes(), &resBody); err != nil {
			t.Fatal(err)
		}

		t.Logf("get user response:\n%+v\n", resBody)
	})
}

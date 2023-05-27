package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"summar/server/cookie"
	"summar/server/handlers"
	"summar/server/stores"
	"summar/server/utils"
	"testing"

	"github.com/google/uuid"
)

func TestRoutes(t *testing.T) {
	mockStore := stores.NewMockStore()
	h := handlers.NewHandlers(mockStore)
	s := NewServer()
	s.MountHandlers(h)

	var sessionToken string
	var bookmarkId uuid.UUID

	t.Run("Test signup route", func(t *testing.T) {
		reqBody := utils.JSONMarshal(&handlers.SignupRequest{
			Email:    "test@test.com",
			Password: "123",
		})

		req, err := http.NewRequest("POST", "/signup", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}

		res := ExecuteRequest(req, s)

		CheckResponseCode(t, http.StatusOK, res.Code)

		resBody := CheckResponseBody[handlers.SignupResponse](t, res.Body.Bytes())
		t.Logf("signup response:\n%+v\n", resBody)
	})

	t.Run("Test login route", func(t *testing.T) {
		reqBody := utils.JSONMarshal(&handlers.LoginRequest{
			Email:    "test@test.com",
			Password: "123",
		})

		req, err := http.NewRequest("POST", "/login", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}

		res := ExecuteRequest(req, s)

		CheckResponseCode(t, http.StatusOK, res.Code)

		resBody := CheckResponseBody[handlers.HandlerResponse](t, res.Body.Bytes())
		t.Logf("login response:\n%+v\n", resBody)

		sessionToken = res.Result().Cookies()[0].Value
		t.Log("session token:", sessionToken)
	})

	t.Run("Test get user route", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/me", nil)
		if err != nil {
			t.Fatal(err)
		}

		res := ExecuteRequest(req, s, sessionToken)
		t.Logf("res: %+v", res)

		CheckResponseCode(t, http.StatusOK, res.Code)

		resBody := CheckResponseBody[handlers.GetUserResponse](t, res.Body.Bytes())
		t.Logf("get user response:\n%+v\n", resBody)
	})

	t.Run("Test create bookmark route", func(t *testing.T) {
		reqBody := utils.JSONMarshal(&handlers.CreateBookmarkRequest{
			Url: "testurl.com",
		})

		req, err := http.NewRequest("POST", "/bookmark", strings.NewReader(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		res := ExecuteRequest(req, s, sessionToken)

		CheckResponseCode(t, http.StatusOK, res.Code)

		resBody := CheckResponseBody[handlers.CreateBookmarkResponse](t, res.Body.Bytes())
		t.Logf("create bookmark response:\n%+v\n", resBody)

		bookmarkId = resBody.Bookmark.Id
	})

	t.Run("Test get bookmark route", func(t *testing.T) {
		url := fmt.Sprintf("/bookmark/%s", bookmarkId.String())

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatal(err)
		}

		res := ExecuteRequest(req, s, sessionToken)
		t.Log("res", res)

		CheckResponseCode(t, http.StatusOK, res.Code)

		resBody := CheckResponseBody[handlers.GetBookmarkResponse](t, res.Body.Bytes())
		t.Logf("get bookmark response:\n%+v\n", resBody)
	})
}

func ExecuteRequest(req *http.Request, s *Server, sessionToken ...string) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	for _, s := range sessionToken {
		cookie.SetSessionTokenCookie(rr, s)
	}

	s.Router.ServeHTTP(rr, req)

	return rr
}

func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func CheckResponseBody[T any](t *testing.T, body []byte) *T {
	var resBody T

	if err := json.Unmarshal(body, &resBody); err != nil {
		t.Fatal(err)
	}

	return &resBody
}

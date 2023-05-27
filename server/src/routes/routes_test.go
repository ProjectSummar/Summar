package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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

	var sessionCookie *http.Cookie
	var bookmarkId uuid.UUID

	t.Run("Test signup route", func(t *testing.T) {
		// format request
		reqBody := utils.JSONMarshal(&handlers.SignupRequest{
			Email:    "test@test.com",
			Password: "123",
		})

		req, err := http.NewRequest("POST", "/signup", strings.NewReader(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		// execute request
		res := ExecuteRequest(req, s)

		// check response
		CheckResponseCode(t, http.StatusOK, res.Code)

		resBody := CheckResponseBody[handlers.SignupResponse](t, res.Body.Bytes())
		t.Logf("signup response:\n%+v\n", resBody)
	})

	t.Run("Test login route", func(t *testing.T) {
		// format request
		reqBody := utils.JSONMarshal(&handlers.LoginRequest{
			Email:    "test@test.com",
			Password: "123",
		})

		req, err := http.NewRequest("POST", "/login", strings.NewReader(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		// execute request
		res := ExecuteRequest(req, s)

		// check response
		CheckResponseCode(t, http.StatusOK, res.Code)

		resBody := CheckResponseBody[handlers.HandlerResponse](t, res.Body.Bytes())
		t.Logf("login response:\n%+v\n", resBody)

		// set external vars
		sessionCookie = res.Result().Cookies()[0]
	})

	t.Run("Test get user route", func(t *testing.T) {
		// format request
		req, err := http.NewRequest("GET", "/me", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.AddCookie(sessionCookie)

		// execute request
		res := ExecuteRequest(req, s)

		// check response
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

		req.AddCookie(sessionCookie)

		res := ExecuteRequest(req, s)

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

		req.AddCookie(sessionCookie)

		res := ExecuteRequest(req, s)

		CheckResponseCode(t, http.StatusOK, res.Code)

		resBody := CheckResponseBody[handlers.GetBookmarkResponse](t, res.Body.Bytes())
		t.Logf("get bookmark response:\n%+v\n", resBody)
	})

	t.Run("Test update bookmark route", func(t *testing.T) {
		// format request
		reqBody := utils.JSONMarshal(&handlers.UpdateBookmarkRequest{
			Url:     utils.StringPtr("updatedurl.com"),
			Summary: utils.StringPtr("updated summary"),
		})

		url := fmt.Sprintf("/bookmark/%s", bookmarkId.String())

		req, err := http.NewRequest("PATCH", url, strings.NewReader(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		req.AddCookie(sessionCookie)

		// execute request
		res := ExecuteRequest(req, s)

		// check response
		CheckResponseCode(t, http.StatusOK, res.Code)

		resBody := CheckResponseBody[handlers.UpdateBookmarkResponse](t, res.Body.Bytes())
		t.Logf("get bookmark response:\n%+v\n", resBody)
	})

	t.Run("Test delete bookmark route", func(t *testing.T) {
		// format request
		url := fmt.Sprintf("/bookmark/%s", bookmarkId.String())

		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			t.Fatal(err)
		}

		req.AddCookie(sessionCookie)

		// execute request
		res := ExecuteRequest(req, s)

		// check response
		CheckResponseCode(t, http.StatusOK, res.Code)

		resBody := CheckResponseBody[handlers.GetBookmarkResponse](t, res.Body.Bytes())
		t.Logf("get bookmark response:\n%+v\n", resBody)
	})
}

func ExecuteRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
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

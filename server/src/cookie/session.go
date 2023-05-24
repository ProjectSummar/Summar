package cookie

import (
	"net/http"
)

const SESSION_TOKEN_COOKIE string = "session-token"

func SetSessionTokenCookie(w http.ResponseWriter, token string) {
	SetCookie(w, SESSION_TOKEN_COOKIE, token)
}

func GetSessionTokenCookie(r *http.Request) (*http.Cookie, error) {
	return GetCookie(r, SESSION_TOKEN_COOKIE)
}

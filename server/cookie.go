package main

import (
	"net/http"
)

func SetCookie(w http.ResponseWriter, name string, value string) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		MaxAge:   3600,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
}

func GetCookie(r *http.Request, name string) (*http.Cookie, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return nil, err
	}

	return cookie, nil
}

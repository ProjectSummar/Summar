package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func ToHttpHandlerFunc(f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, HandlerResponse{
				IsOk: false,
				Msg:  err.Error(),
			})
		}
	}
}

type (
	Middleware     func(next http.Handler) http.Handler
	MiddlewareFunc func(w http.ResponseWriter, r *http.Request) (context.Context, error)
)

func ToMiddleware(f MiddlewareFunc) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, err := f(w, r)
			if err != nil {
				WriteJSON(w, http.StatusBadRequest, HandlerResponse{
					IsOk: false,
					Msg:  err.Error(),
				})
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func ReadJSON(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

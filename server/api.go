package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type APIServer struct {
	Address string
}

func NewAPIServer(address string) *APIServer {
	return &APIServer{
		Address: address,
	}
}

func (s *APIServer) Run() {
	router := httprouter.New()
	// TODO: register routes
	// router.GET("/example", ToRouterHandleFunc(ExampleHandler))
	log.Fatal(http.ListenAndServe(s.Address, router))
}

// TODO: routes here
// func (s *APIServer) ExampleHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
// 	return WriteJSON(w, http.StatusOK, res)
// }

type APIFunc func(http.ResponseWriter, *http.Request, httprouter.Params) error

type APIError struct {
	Error string `json:"error"`
}

func ToRouterHandleFunc(f APIFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := f(w, r, p); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{
				Error: err.Error(),
			})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

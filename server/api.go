package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type APIResponse struct {
	IsOk bool   `json:"isOk"`
	Msg  string `json:"msg"`
}

type APIServer struct {
	Address string
}

func NewAPIServer(address string) *APIServer {
	return &APIServer{
		Address: address,
	}
}

func (s *APIServer) Run() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/login", ToHttpHandlerFunc(s.LoginHandler))
	router.Post("/signup", ToHttpHandlerFunc(s.SignupHandler))
	router.Get("/me", ToHttpHandlerFunc(s.GetUserHandler))

	router.Mount("/bookmark", s.BookmarkRouter())

	log.Println("server running on port", s.Address)
	log.Fatal(http.ListenAndServe(":"+s.Address, router))
}

func (s *APIServer) BookmarkRouter() chi.Router {
	bookmarkRouter := chi.NewRouter()

	bookmarkRouter.Post("/create", ToHttpHandlerFunc(s.CreateBookmarkHandler))
	bookmarkRouter.Get("/get", ToHttpHandlerFunc(s.GetBookmarkHandler))
	bookmarkRouter.Post("/update", ToHttpHandlerFunc(s.UpdateBookmarkHandler))
	bookmarkRouter.Post("/delete", ToHttpHandlerFunc(s.DeleteBookmarkHandler))
	bookmarkRouter.Post("/summarise", ToHttpHandlerFunc(s.SummariseBookmarkHandler))

	return bookmarkRouter
}

// TODO: auth middleware to get session token from cookie and validate

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	APIResponse
}

func (s *APIServer) LoginHandler(w http.ResponseWriter, r *http.Request) error {
	// parse input JSON { email, password }
	var loginRequest LoginRequest
	ReadJSON(r, &loginRequest)

	// validate credentials
	// create a session
	// return session token to be stored as cookie

	WriteJSON(w, http.StatusOK, LoginResponse{
		APIResponse: APIResponse{
			IsOk: true,
			Msg:  "logged in successfully",
		},
	})

	return nil
}

func (s *APIServer) SignupHandler(w http.ResponseWriter, r *http.Request) error {
	// parse input JSON { email, password }
	// validate credentials
	// hash password
	// create user
	// create a session
	// return session token to be stored as cookie
	return nil
}

func (s *APIServer) GetUserHandler(w http.ResponseWriter, r *http.Request) error {
	// get session token in cookie
	// validate session token
	// get user by userId associated to the session
	// get bookmarks by userId
	// return user and bookmarks
	return nil
}

func (s *APIServer) CreateBookmarkHandler(w http.ResponseWriter, r *http.Request) error {
	// get session token in cookie
	// validate session token
	// get user by userId associated to the session
	// parse input JSON { url }
	// create bookmark on userId
	// return status and created bookmark
	return nil
}

func (s *APIServer) GetBookmarkHandler(w http.ResponseWriter, r *http.Request) error {
	// get session token in cookie
	// validate session token
	// get user by userId associated to the session
	// parse input JSON { bookmarkId }
	// get bookmark by bookmarkId
	// validate bookmark's userId and session userId
	// return status and bookmark
	return nil
}

func (s *APIServer) UpdateBookmarkHandler(w http.ResponseWriter, r *http.Request) error {
	// get session token in cookie
	// validate session token
	// get user by userId associated to the session
	// parse input JSON { partialBookmark }
	// get bookmark by bookmarkId
	// validate bookmark's userId and session userId
	// update bookmark with partialBookmark
	// return status and updated bookmark
	return nil
}

func (s *APIServer) DeleteBookmarkHandler(w http.ResponseWriter, r *http.Request) error {
	// get session token in cookie
	// validate session token
	// get user by userId associated to the session
	// parse input JSON { bookmarkId }
	// get bookmark by bookmarkId
	// validate bookmark's userId and session userId
	// delete bookmark by bookmarkId
	// return status and deleted bookmark
	return nil
}

func (s *APIServer) SummariseBookmarkHandler(w http.ResponseWriter, r *http.Request) error {
	// parse input JSON { bookmarkId }
	// get bookmark by bookmarkId
	// scrape bookmark for main text
	// send main text to OpenAI API for summarisation
	// update bookmark with summary
	// return status and summarised bookmark
	return nil
}

type APIFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	APIResponse
}

func ToHttpHandlerFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{
				APIResponse: APIResponse{
					IsOk: false,
					Msg:  err.Error(),
				},
			})
		}
	}
}

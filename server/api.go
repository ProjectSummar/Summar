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
	Store   Store
}

func NewAPIServer(address string, store Store) *APIServer {
	return &APIServer{
		Address: address,
		Store:   store,
	}
}

func (s *APIServer) Run() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/login", ToHttpHandlerFunc(s.LoginHandler))
	router.Post("/signup", ToHttpHandlerFunc(s.SignupHandler))
	router.Get("/me", ToHttpHandlerFunc(s.GetUserHandler))

	bookmarkRouter := chi.NewRouter()

	bookmarkRouter.Post("/create", ToHttpHandlerFunc(s.CreateBookmarkHandler))
	bookmarkRouter.Get("/get", ToHttpHandlerFunc(s.GetBookmarkHandler))
	bookmarkRouter.Post("/update", ToHttpHandlerFunc(s.UpdateBookmarkHandler))
	bookmarkRouter.Post("/delete", ToHttpHandlerFunc(s.DeleteBookmarkHandler))
	bookmarkRouter.Post("/summarise", ToHttpHandlerFunc(s.SummariseBookmarkHandler))

	router.Mount("/bookmark", bookmarkRouter)

	log.Println("Server running on port", s.Address)
	log.Fatal(http.ListenAndServe(":"+s.Address, router))
}

// TODO: auth middleware to get session token from cookie and validate
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ctx := context.WithValue(r.Context(), "user", "123")
		// next.ServeHTTP(w, r.WithContext(ctx))
		next.ServeHTTP(w, r)
	})
}

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
	user, err := s.Store.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return err
	}

	if err := ComparePasswordToHash(user.PasswordHash, loginRequest.Password); err != nil {
		return err
	}

	// create a session
	sessionToken, err := CreateSessionToken()
	if err != nil {
		return err
	}

	session := NewSession(sessionToken, user.ID)

	if err := s.Store.CreateSession(&session); err != nil {
		return err
	}

	// return session token to be stored as cookie
	SetSessionTokenCookie(w, sessionToken)

	WriteJSON(w, http.StatusOK, LoginResponse{
		APIResponse: APIResponse{
			IsOk: true,
			Msg:  "Logged in successfully",
		},
	})

	return nil
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	APIResponse
}

func (s *APIServer) SignupHandler(w http.ResponseWriter, r *http.Request) error {
	// parse input JSON { email, password }
	var signupRequest SignupRequest
	ReadJSON(r, &signupRequest)

	// hash password
	hash, err := HashPassword(signupRequest.Password)
	if err != nil {
		return err
	}

	// create user
	user := NewUser(signupRequest.Email, hash)
	if err := s.Store.CreateUser(&user); err != nil {
		return err
	}

	// create a session
	sessionToken, err := CreateSessionToken()
	if err != nil {
		return err
	}

	session := NewSession(sessionToken, user.ID)

	if err := s.Store.CreateSession(&session); err != nil {
		return err
	}

	// return session token to be stored as cookie
	SetSessionTokenCookie(w, sessionToken)

	WriteJSON(w, http.StatusOK, SignupResponse{
		APIResponse: APIResponse{
			IsOk: true,
			Msg:  "Signed up successfully",
		},
	})

	return nil
}

type GetUserResponse struct {
	APIResponse
	User      *User      `json:"user"`
	Bookmarks []Bookmark `json:"bookmarks"`
}

func (s *APIServer) GetUserHandler(w http.ResponseWriter, r *http.Request) error {
	// get session token in cookie
	cookie, err := GetSessionTokenCookie(r)
	if err != nil {
		return err
	}

	sessionToken := cookie.Value

	// validate session expiry
	session, err := s.Store.GetSessionByToken(sessionToken)
	if err != nil {
		return err
	}

	if err := ValidateSessionExpiry(session); err != nil {
		return err
	}

	// get associated user
	user, err := s.Store.GetUserById(session.UserID)
	if err != nil {
		return err
	}

	// TODO: get bookmarks by userId
	// bookmarks, err := s.Db.GetBookmarksByUserId(user.ID)

	// return user and bookmarks
	WriteJSON(w, http.StatusOK, GetUserResponse{
		APIResponse: APIResponse{
			IsOk: true,
			Msg:  "Got user successfully",
		},
		User: user,
		// Bookmarks: bookmarks,
	})

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

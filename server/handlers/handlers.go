package handlers

import (
	"encoding/json"
	"net/http"
	"summar/server/cookie"
	"summar/server/password"
	"summar/server/stores"
	"summar/server/types"
)

type Handlers struct {
	Store stores.Store
}

func NewHandlers(store stores.Store) *Handlers {
	return &Handlers{
		Store: store,
	}
}

type HandlerResponse struct {
	IsOk bool   `json:"isOk"`
	Msg  string `json:"msg"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	HandlerResponse
}

func (h *Handlers) LoginHandler(w http.ResponseWriter, r *http.Request) error {
	// parse input JSON { email, password }
	var loginRequest LoginRequest
	ReadJSON(r, &loginRequest)

	// validate credentials
	user, err := h.Store.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return err
	}

	if err := password.CompareHashToPassword(user.PasswordHash, loginRequest.Password); err != nil {
		return err
	}

	// create a session
	session := types.NewSession(user.Id)

	if err := h.Store.CreateSession(&session); err != nil {
		return err
	}

	// return session token to be stored as cookie
	cookie.SetSessionTokenCookie(w, session.Token)

	WriteJSON(w, http.StatusOK, LoginResponse{
		HandlerResponse: HandlerResponse{
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
	HandlerResponse
}

func (h *Handlers) SignupHandler(w http.ResponseWriter, r *http.Request) error {
	// parse input JSON { email, password }
	var signupRequest SignupRequest
	ReadJSON(r, &signupRequest)

	// hash password
	hash, err := password.HashPassword(signupRequest.Password)
	if err != nil {
		return err
	}

	// create user
	user := types.NewUser(signupRequest.Email, hash)
	if err := h.Store.CreateUser(&user); err != nil {
		return err
	}

	WriteJSON(w, http.StatusOK, SignupResponse{
		HandlerResponse: HandlerResponse{
			IsOk: true,
			Msg:  "Signed up successfully",
		},
	})

	return nil
}

type GetUserResponse struct {
	HandlerResponse
	User      *types.User       `json:"user"`
	Bookmarks []*types.Bookmark `json:"bookmarks"`
}

func (h *Handlers) GetUserHandler(w http.ResponseWriter, r *http.Request) error {
	// get session token in cookie
	cookie, err := cookie.GetSessionTokenCookie(r)
	if err != nil {
		return err
	}

	sessionToken := cookie.Value

	// validate session expiry
	session, err := h.Store.GetSession(sessionToken)
	if err != nil {
		return err
	}

	if err := types.VerifySessionExpiry(session); err != nil {
		return err
	}

	// get associated user
	user, err := h.Store.GetUser(session.UserId)
	if err != nil {
		return err
	}

	// TODO: get bookmarks by userId
	// bookmarks, err := s.Db.GetBookmarksByUserId(user.ID)

	// return user and bookmarks
	WriteJSON(w, http.StatusOK, GetUserResponse{
		HandlerResponse: HandlerResponse{
			IsOk: true,
			Msg:  "Got user successfully",
		},
		User: user,
		// Bookmarks: bookmarks,
	})

	return nil
}

func (h *Handlers) CreateBookmarkHandler(w http.ResponseWriter, r *http.Request) error {
	// get session token in cookie
	// validate session token
	// get user by userId associated to the session
	// parse input JSON { url }
	// create bookmark on userId
	// return status and created bookmark
	return nil
}

func (h *Handlers) GetBookmarkHandler(w http.ResponseWriter, r *http.Request) error {
	// get session token in cookie
	// validate session token
	// get user by userId associated to the session
	// parse input JSON { bookmarkId }
	// get bookmark by bookmarkId
	// validate bookmark's userId and session userId
	// return status and bookmark
	return nil
}

func (h *Handlers) UpdateBookmarkHandler(w http.ResponseWriter, r *http.Request) error {
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

func (h *Handlers) DeleteBookmarkHandler(w http.ResponseWriter, r *http.Request) error {
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

func (h *Handlers) SummariseBookmarkHandler(w http.ResponseWriter, r *http.Request) error {
	// parse input JSON { bookmarkId }
	// get bookmark by bookmarkId
	// scrape bookmark for main text
	// send main text to OpenAI API for summarisation
	// update bookmark with summary
	// return status and summarised bookmark
	return nil
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func ToHttpHandler(f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, HandlerResponse{
				IsOk: false,
				Msg:  err.Error(),
			})
		}
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

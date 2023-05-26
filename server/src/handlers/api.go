package handlers

import (
	"net/http"
	"summar/server/cookie"
	"summar/server/password"
	"summar/server/types"

	"github.com/google/uuid"
)

type HandlerResponse struct {
	IsOk bool   `json:"isOk"`
	Msg  string `json:"msg"`
}

type (
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

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

	if err := h.Store.CreateSession(session); err != nil {
		return err
	}

	// return session token to be stored as cookie
	cookie.SetSessionTokenCookie(w, session.Token)

	return WriteJSON(w, http.StatusOK, &HandlerResponse{
		IsOk: true,
		Msg:  "Logged in successfully",
	})
}

type (
	SignupRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

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

	if err := h.Store.CreateUser(user); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, &HandlerResponse{
		IsOk: true,
		Msg:  "Signed up successfully",
	})
}

type (
	GetUserResponse struct {
		HandlerResponse
		User      *types.User       `json:"user"`
		Bookmarks []*types.Bookmark `json:"bookmarks"`
	}
)

func (h *Handlers) GetUserHandler(w http.ResponseWriter, r *http.Request) error {
	// get userId from auth middleware context
	userId := r.Context().Value("userId").(uuid.UUID)

	// get associated user
	user, err := h.Store.GetUser(userId)
	if err != nil {
		return err
	}

	// TODO: get bookmarks by userId
	// bookmarks, err := s.Db.GetBookmarksByUserId(user.ID)

	// return user and bookmarks
	return WriteJSON(w, http.StatusOK, &GetUserResponse{
		HandlerResponse: HandlerResponse{
			IsOk: true,
			Msg:  "Got user successfully",
		},
		User: user,
		// Bookmarks: bookmarks,
	})
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

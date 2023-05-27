package handlers

import (
	"fmt"
	"net/http"
	"summar/server/cookie"
	"summar/server/password"
	"summar/server/types"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type HandlerResponse struct {
	Ok  bool   `json:"ok"`
	Msg string `json:"msg"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handlers) LoginHandler(w http.ResponseWriter, r *http.Request) error {
	// parse input JSON { email, password }
	var req LoginRequest
	ReadJSON(r, &req)

	// validate credentials
	user, err := h.Store.GetUserByEmail(req.Email)
	if err != nil {
		return err
	}

	if err := password.CompareHashToPassword(user.PasswordHash, req.Password); err != nil {
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
		Ok:  true,
		Msg: "Logged in successfully",
	})
}

type (
	SignupRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignupResponse struct {
		HandlerResponse
		User *types.User `json:"user"`
	}
)

func (h *Handlers) SignupHandler(w http.ResponseWriter, r *http.Request) error {
	// parse input JSON { email, password }
	var req SignupRequest
	ReadJSON(r, &req)

	// hash password
	hash, err := password.HashPassword(req.Password)
	if err != nil {
		return err
	}

	// create user
	user, err := types.NewUser(req.Email, hash)
	if err != nil {
		return err
	}

	if err := h.Store.CreateUser(user); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, &SignupResponse{
		HandlerResponse: HandlerResponse{
			Ok:  true,
			Msg: "Signed up successfully",
		},
		User: user,
	})
}

type GetUserResponse struct {
	HandlerResponse
	User      *types.User       `json:"user"`
	Bookmarks []*types.Bookmark `json:"bookmarks"`
}

func (h *Handlers) GetUserHandler(w http.ResponseWriter, r *http.Request) error {
	// get userId from auth middleware context
	userId := r.Context().Value("userId").(uuid.UUID)

	// get associated user
	user, err := h.Store.GetUser(userId)
	if err != nil {
		return err
	}

	// get bookmarks by userId
	bookmarks, err := h.Store.GetBookmarksByUserId(user.Id)
	if err != nil {
		return err
	}

	// return user and bookmarks
	return WriteJSON(w, http.StatusOK, &GetUserResponse{
		HandlerResponse: HandlerResponse{
			Ok:  true,
			Msg: "Got user successfully",
		},
		User:      user,
		Bookmarks: bookmarks,
	})
}

type (
	CreateBookmarkRequest struct {
		Url string `json:"url"`
	}

	CreateBookmarkResponse struct {
		HandlerResponse
		Bookmark *types.Bookmark `json:"bookmark"`
	}
)

func (h *Handlers) CreateBookmarkHandler(w http.ResponseWriter, r *http.Request) error {
	// get userId from auth middleware context
	userId := r.Context().Value("userId").(uuid.UUID)

	// parse input JSON { url }
	var req CreateBookmarkRequest
	if err := ReadJSON(r, &req); err != nil {
		return err
	}

	// create bookmark on userId
	bookmark, err := types.NewBookmark(userId, req.Url)
	if err != nil {
		return err
	}

	if err := h.Store.CreateBookmark(bookmark); err != nil {
		return err
	}

	// return status and created bookmark
	return WriteJSON(w, http.StatusOK, &CreateBookmarkResponse{
		HandlerResponse: HandlerResponse{
			Ok:  true,
			Msg: "Bookmark created successfully",
		},
		Bookmark: bookmark,
	})
}

type GetBookmarkResponse struct {
	HandlerResponse
	Bookmark *types.Bookmark `json:"bookmark"`
}

func (h *Handlers) GetBookmarkHandler(w http.ResponseWriter, r *http.Request) error {
	// get userId from auth middleware context
	userId := r.Context().Value("userId").(uuid.UUID)

	// get id from url param and parse to uuid
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	// get bookmark by bookmarkId
	bookmark, err := h.Store.GetBookmark(id)
	if err != nil {
		return err
	}

	// validate bookmark's userId and session userId
	if userId != bookmark.UserId {
		return fmt.Errorf("Unauthorised to view this bookmark")
	}

	// return status and bookmark
	return WriteJSON(w, http.StatusOK, &GetBookmarkResponse{
		HandlerResponse: HandlerResponse{
			Ok:  true,
			Msg: "Got bookmark successfully",
		},
		Bookmark: bookmark,
	})
}

type (
	UpdateBookmarkRequest struct {
		Url     *string `json:"url,omitempty"`
		Summary *string `json:"summary,omitempty"`
	}

	UpdateBookmarkResponse struct {
		HandlerResponse
		Bookmark *types.Bookmark `json:"bookmark"`
	}
)

func (h *Handlers) UpdateBookmarkHandler(w http.ResponseWriter, r *http.Request) error {
	// get userId from auth middleware context
	userId := r.Context().Value("userId").(uuid.UUID)

	// get id from url param and parse to uuid
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	// parse input JSON { partialBookmark }
	var req UpdateBookmarkRequest
	if err := ReadJSON(r, &req); err != nil {
		return err
	}

	// get bookmark by bookmarkId
	bookmark, err := h.Store.GetBookmark(id)
	if err != nil {
		return err
	}

	// validate bookmark's userId and session userId
	if userId != bookmark.UserId {
		return fmt.Errorf("Unauthorised to view this bookmark")
	}

	// update bookmark with partialBookmark
	if req.Url != nil {
		bookmark.Url = *req.Url
	}

	if req.Summary != nil {
		bookmark.Summary = *req.Summary
	}

	if err := h.Store.UpdateBookmark(bookmark); err != nil {
		return err
	}

	// return status and updated bookmark
	return WriteJSON(w, http.StatusOK, &UpdateBookmarkResponse{
		HandlerResponse: HandlerResponse{
			Ok:  true,
			Msg: "Bookmark updated successfully",
		},
		Bookmark: bookmark,
	})
}

type DeleteBookmarkResponse struct {
	HandlerResponse
	Bookmark *types.Bookmark `json:"bookmark"`
}

func (h *Handlers) DeleteBookmarkHandler(w http.ResponseWriter, r *http.Request) error {
	// get userId from auth middleware context
	userId := r.Context().Value("userId").(uuid.UUID)

	// get id from url param and parse to uuid
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	// get bookmark by bookmarkId
	bookmark, err := h.Store.GetBookmark(id)
	if err != nil {
		return err
	}

	// validate bookmark's userId and session userId
	if userId != bookmark.UserId {
		return fmt.Errorf("Unauthorised to view this bookmark")
	}

	// delete bookmark by bookmarkId
	if err := h.Store.DeleteBookmark(id); err != nil {
		return err
	}

	// return status and deleted bookmark
	return WriteJSON(w, http.StatusOK, &DeleteBookmarkResponse{
		HandlerResponse: HandlerResponse{
			Ok:  true,
			Msg: "Bookmark deleted successfully",
		},
		Bookmark: bookmark,
	})
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

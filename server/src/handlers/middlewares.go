package handlers

import (
	"context"
	"fmt"
	"net/http"
	"summar/server/cookie"
	"summar/server/types"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *Handlers) AuthMiddlewareFunc(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	// get session token in cookie
	cookie, err := cookie.GetSessionTokenCookie(r)
	if err != nil {
		return nil, err
	}

	sessionToken := cookie.Value

	// get session from store
	session, err := h.Store.GetSession(sessionToken)
	if err != nil {
		return nil, err
	}

	// validate session expiry
	if err := types.ValidateSession(session); err != nil {
		return nil, err
	}

	ctx := context.WithValue(r.Context(), "userId", session.UserId)

	return ctx, nil
}

func (h *Handlers) BookmarkMiddlewareFunc(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	// get userId from auth middleware context
	userId := r.Context().Value("userId").(uuid.UUID)

	// get bookmark bookmarkId from url param and parse to uuid
	bookmarkId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return nil, err
	}

	// get bookmark by bookmarkId
	bookmark, err := h.Store.GetBookmark(bookmarkId)
	if err != nil {
		return nil, err
	}

	// validate bookmark's userId and session userId
	if userId != bookmark.UserId {
		return nil, fmt.Errorf("Unauthorised to view/modify this bookmark")
	}

	// set bookmark in context
	ctx := context.WithValue(r.Context(), "bookmark", bookmark)

	return ctx, nil
}

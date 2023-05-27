package handlers

import (
	"context"
	"net/http"
	"summar/server/cookie"
	"summar/server/types"
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

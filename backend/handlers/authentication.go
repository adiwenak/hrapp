package handlers

import (
	"context"

	"github.com/adiwenak/hrapp/ent/user"
	"github.com/adiwenak/hrapp/middlewares"
	"github.com/adiwenak/hrapp/server"
)

// Login
// (POST /login)
func (h *HRCore) Login(ctx context.Context, request server.LoginRequestObject) (server.LoginResponseObject, error) {
	val, err := h.Client.User.Query().Where(user.Email(request.Body.Email)).Only(ctx)
	if err != nil {
		return nil, err
	}

	if val.Password != request.Body.Password {
		return nil, err
	}

	sessionCookie, err := middlewares.CreateCookieAccessToken(val.Email)

	if err != nil {
		return nil, err
	}

	return server.Login200Response{
		Headers: server.Login200ResponseHeaders{
			SetCookie: sessionCookie,
		},
	}, nil
}

// Logout
// (POST /logout)
func (h *HRCore) Logout(ctx context.Context, request server.LogoutRequestObject) (server.LogoutResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

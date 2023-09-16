package handlers

import (
	"context"
	"time"

	"github.com/adiwenak/hrapp/ent/user"
	"github.com/adiwenak/hrapp/server"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key")

// Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func createSessionCookie(username string) (string, error) {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

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

	sessionCookie, err := createSessionCookie(val.Email)

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

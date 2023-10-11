package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key")

// Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateCookieAccessToken(username string) (string, error) {
	accessToken, err := CreateAccessToken(username)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("access_token=%s", accessToken), nil
}

func CreateAccessToken(username string) (string, error) {
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

func getUsernameFromRequest(request *http.Request) (string, error) {
	c, err := request.Cookie("access_token")

	if err != nil {
		return "", err
	}

	accessToken := c.Value

	claims := &Claims{}

	_, err = jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})

	username := claims.Username

	if err != nil {
		return "", err
	}

	return username, nil
}

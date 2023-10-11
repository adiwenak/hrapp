package middlewares

import (
	"context"
	"fmt"

	"github.com/adiwenak/hrapp/server"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gofiber/fiber/v2"
	middleware "github.com/oapi-codegen/fiber-middleware"
)

type ctxKeyUsername struct{}

func CreateOapiMiddlewares() (fiber.Handler, error) {
	spec, err := server.GetSwagger()
	if err != nil {
		return nil, err
	}

	validator := middleware.OapiRequestValidatorWithOptions(spec, &middleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: newAuthenticator(),
		},
	})

	return validator, nil
}

func GetUsernameFromUserContext(ctx context.Context) (string, bool) {
	val, err := ctx.Value(ctxKeyUsername{}).(string)
	return val, err
}

func newAuthenticator() openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		if input.SecuritySchemeName != "cookieAuth" {
			return fmt.Errorf("security scheme %s != 'cookieAuth'", input.SecuritySchemeName)
		}

		username, err := getUsernameFromRequest(input.RequestValidationInput.Request)
		if err != nil {
			return err
		}

		fiberContext := middleware.GetFiberContext(ctx)
		userContext := fiberContext.UserContext()
		newContext := context.WithValue(userContext, ctxKeyUsername{}, username)

		fiberContext.SetUserContext(newContext)

		if err != nil {
			return err
		}

		return nil
	}
}

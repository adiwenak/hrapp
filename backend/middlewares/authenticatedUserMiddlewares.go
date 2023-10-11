package middlewares

import (
	"context"

	"github.com/adiwenak/hrapp/ent"
	"github.com/adiwenak/hrapp/ent/user"
	"github.com/gofiber/fiber/v2"
)

type ctxKeyUser struct{}

func GetUserFromUserContext(ctx context.Context) (*ent.User, bool) {
	val, err := ctx.Value(ctxKeyUser{}).(*ent.User)
	return val, err
}

func CreateAuthenticatedUserMiddlewares(dbClient *ent.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userContext := ctx.UserContext()
		username, ok := GetUsernameFromUserContext(userContext)

		if !ok {
			return ctx.Next()
		}

		user, err := dbClient.User.Query().Where(user.Email(username)).Only(userContext)

		if err != nil {
			return err
		}

		newContext := context.WithValue(userContext, ctxKeyUser{}, user)

		ctx.SetUserContext(newContext)

		return ctx.Next()
	}
}

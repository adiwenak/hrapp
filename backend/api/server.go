package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/adiwenak/hrapp/ent"
	"github.com/gofiber/fiber/v2"
)

type ServerInterfaceImpl struct {
	Client *ent.Client
}

func (serv *ServerInterfaceImpl) CreateUser(c *fiber.Ctx) error {
	newUser := new(NewUser)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	_, err := serv.Client.User.Create().SetFirstName(*newUser.FirstName).SetLastName(*newUser.LastName).Save(c.Context())

	if err != nil {
		fmt.Errorf("failed creating user: %w", err)
		c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(newUser)
}

func (serv *ServerInterfaceImpl) GetUserProfile(c *fiber.Ctx, userid int64) error {
	return c.SendString("User Profile")
}

func (serv *ServerInterfaceImpl) UserCheckIn(c *fiber.Ctx, userid int64) error {
	return c.SendString("User Checkin")
}

func (serv *ServerInterfaceImpl) UserCheckOut(c *fiber.Ctx, userid int64) error {
	return c.SendString("User Checkout")
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetFirstName("Ben").
		SetLastName("Tony").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

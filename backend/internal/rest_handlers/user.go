package rest_handlers

import (
	"fmt"
	"net/http"

	"github.com/adiwenak/hrapp/api"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (serv *ServerInterfaceImpl) CreateUser(c *fiber.Ctx) error {
	newUser := new(api.NewUser)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	if err := serv.Validate.Struct(newUser); err != nil {
		var validationErrors = err.(validator.ValidationErrors)
		for _, err := range validationErrors {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		return c.Status(http.StatusBadRequest).JSON(validationErrors)
	}

	_, err := serv.Client.User.Create().
		SetFirstName(newUser.FirstName).
		SetLastName(newUser.LastName).
		Save(c.Context())

	if err != nil {
		fmt.Errorf("failed creating user: %w", err)
		c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(newUser)
}

func (serv *ServerInterfaceImpl) GetUserProfile(c *fiber.Ctx, userid string) error {
	return c.SendString("User Profile")
}

func (serv *ServerInterfaceImpl) UserCheckIn(c *fiber.Ctx, userid string) error {
	return c.SendString("User Checkin")
}

func (serv *ServerInterfaceImpl) UserCheckOut(c *fiber.Ctx, userid string) error {
	return c.SendString("User Checkout")
}

func (serv *ServerInterfaceImpl) GenerateVerificationCode(c *fiber.Ctx, userid string) error {
	return c.SendString("Generate Verification Code")
}

func (serv *ServerInterfaceImpl) ChangeUserPassword(c *fiber.Ctx, userid string) error {
	return c.SendString("Change User Password")
}
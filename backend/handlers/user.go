package handlers

import (
	"context"
	"fmt"

	"github.com/adiwenak/hrapp/middlewares"
	"github.com/adiwenak/hrapp/server"
	"github.com/adiwenak/hrapp/utils"
	"github.com/go-playground/validator/v10"
)

// Get current user profile
// (GET /me/profile)
func (h *HRCore) GetCurrentUserProfile(ctx context.Context, request server.GetCurrentUserProfileRequestObject) (server.GetCurrentUserProfileResponseObject, error) {
	usr, ok := middlewares.GetUserFromUserContext(ctx)

	if !ok {
		return nil, fmt.Errorf("unable to get user")
	}

	return server.GetCurrentUserProfile200JSONResponse{
		FirstName:    usr.FirstName,
		LastName:     usr.LastName,
		Email:        usr.Email,
		MobileNumber: usr.MobileNumber,
	}, nil

}

// Get users
// (GET /users)
func (h *HRCore) GetUsers(ctx context.Context, request server.GetUsersRequestObject) (server.GetUsersResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Create user
// (POST /users)
func (h *HRCore) CreateUser(ctx context.Context, request server.CreateUserRequestObject) (server.CreateUserResponseObject, error) {
	newUser := request.Body
	if err := h.Validate.Struct(newUser); err != nil {
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

		return nil, err
	}

	createdUser, err := h.Client.User.Create().
		SetFirstName(newUser.FirstName).
		SetLastName(newUser.LastName).
		SetEmail(newUser.Email).
		SetMobileNumber(newUser.MobileNumber).
		SetUsername(newUser.Username).
		SetPassword(utils.RandomPassword(8)).
		SetNeedPasswordReset(true).
		SetOrganisationID(request.Body.OrgId).
		Save(ctx)

	if err != nil {
		fmt.Errorf("failed creating user: %w", err)
		return nil, err
	}

	return server.CreateUser200JSONResponse{
		Email:             createdUser.Email,
		FirstName:         createdUser.FirstName,
		Id:                createdUser.ID,
		LastName:          createdUser.LastName,
		MobileNumber:      createdUser.MobileNumber,
		TemporaryPassword: createdUser.Password,
	}, err
}

package handlers

import (
	"context"
	"fmt"

	"github.com/adiwenak/hrapp/server"
	"github.com/go-playground/validator/v10"
)

// Create organisation
// (POST /organisations)
func (h *HRCore) CreateOrganisation(ctx context.Context, request server.CreateOrganisationRequestObject) (server.CreateOrganisationResponseObject, error) {
	newOrg := request.Body
	if err := h.Validate.Struct(newOrg); err != nil {
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

	createdOrg, err := h.Client.Organisation.Create().SetName(newOrg.Name).Save(ctx)

	if err != nil {
		return nil, err
	}

	return server.CreateOrganisation200JSONResponse{
		Id:   createdOrg.ID,
		Name: createdOrg.Name,
	}, nil
}

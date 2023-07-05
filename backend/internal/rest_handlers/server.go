package rest_handlers

import (
	"github.com/adiwenak/hrapp/ent"
	"github.com/go-playground/validator/v10"
)

type ServerInterfaceImpl struct {
	Client   *ent.Client
	Validate *validator.Validate
}

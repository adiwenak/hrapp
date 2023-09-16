package handlers

import (
	"github.com/adiwenak/hrapp/ent"
	"github.com/adiwenak/hrapp/server"
	"github.com/go-playground/validator/v10"
)

var _ server.StrictServerInterface = (*HRCore)(nil)

type HRCore struct {
	Client   *ent.Client
	Validate *validator.Validate
}

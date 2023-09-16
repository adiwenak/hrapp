package handlers

import (
	"context"

	"github.com/adiwenak/hrapp/server"
)

// User check in
// (POST /me/checkIn)
func (h *HRCore) UserCheckIn(ctx context.Context, request server.UserCheckInRequestObject) (server.UserCheckInResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// User check out
// (POST /me/checkOut)
func (h *HRCore) UserCheckOut(ctx context.Context, request server.UserCheckOutRequestObject) (server.UserCheckOutResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

package handlers

import (
	"context"

	"github.com/adiwenak/hrapp/server"
)

// Change user password
// (PUT /me/selfservice/changepassword)
func (mp HRCore) ChangeUserPassword(ctx context.Context, request server.ChangeUserPasswordRequestObject) (server.ChangeUserPasswordResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Self service verification code
// (PUT /me/selfservice/changepassword/verificationCode)
func (h *HRCore) GenerateVerificationCode(ctx context.Context, request server.GenerateVerificationCodeRequestObject) (server.GenerateVerificationCodeResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

package interfaces

import (
	"context"

	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
)

type OtpUseCase interface {
	SendOTP(ctx context.Context, mobno request.OTPreq) (string, error)
	VerifyOTP(ctx context.Context, userData request.Otpverifier) error
}

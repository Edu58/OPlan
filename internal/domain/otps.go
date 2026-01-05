package domain

import (
	"context"

	"github.com/Edu58/Oplan/internal/database/sqlc"
)

type OTPRepository interface {
	CreateOTP(ctx context.Context, arg sqlc.CreateOTPParams) (sqlc.OtpStore, error)
	GetOTP(ctx context.Context, identifier string) (sqlc.OtpStore, error)
	UpdateOTP(ctx context.Context, arg sqlc.UpdateOTPParams) (sqlc.OtpStore, error)
	DeleteOTP(ctx context.Context, identifier string) error
}

package services

import (
	"context"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/Edu58/Oplan/internal/domain"
	"github.com/Edu58/Oplan/pkg/logger"
)

type OTPService struct {
	repo   domain.OTPRepository
	logger logger.Logger
}

func NewOTPHandler(repo domain.OTPRepository, logger logger.Logger) *OTPService {
	return &OTPService{repo, logger}
}

func (u *OTPService) GetOTP(ctx context.Context, identifier string) (sqlc.OtpStore, error) {
	return u.repo.GetOTP(ctx, identifier)
}

func (u *OTPService) CreateOTP(ctx context.Context, arg sqlc.CreateOTPParams) (sqlc.OtpStore, error) {
	return u.repo.CreateOTP(ctx, arg)
}
func (u *OTPService) UpdateOTP(ctx context.Context, arg sqlc.UpdateOTPParams) (sqlc.OtpStore, error) {
	return u.repo.UpdateOTP(ctx, arg)
}

func (u *OTPService) DeleteOTP(ctx context.Context, identifier string) error {
	return u.repo.DeleteOTP(ctx, identifier)
}

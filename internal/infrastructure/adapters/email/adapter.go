package adapters

import (
	"context"

	"github.com/Locky-Inc/locky/internal/domain"
)

type SmtpClient interface {
	SendEmail(ctx context.Context, creds *domain.SmtpCredentials, email *domain.Email) error
}

type Adapter struct {
	smtpClient SmtpClient
}

func New(smtpClinet SmtpClient) *Adapter {
	return &Adapter{
		smtpClient: smtpClinet,
	}
}

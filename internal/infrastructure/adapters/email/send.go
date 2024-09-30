package adapters

import (
	"context"
	"github.com/Locky-Inc/locky/internal/domain"
)

func (a *Adapter) SendEmail(ctx context.Context, creds *domain.EmailProviderCredentials, email *domain.Email) error {
	switch creds.Type {
	case domain.EmailProviderTypeSmtp:
		if creds.SMTP == nil {
			return ErrInvalidCredentials
		}
		// TODO надо понять какие SMTP ошибки являются понятными терминальными и явно их отмечать
		return a.smtpClient.SendEmail(ctx, creds.SMTP, email)
	default:
		// TODO добавить поддержку fake provider
		return ErrUnsupportedCredentials
	}
}

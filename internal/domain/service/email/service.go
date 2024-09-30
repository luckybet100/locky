package email

import "github.com/Locky-Inc/locky/internal/domain"

type Service struct {
	creds []*domain.EmailProviderCredentials
}

func New(creds []*domain.EmailProviderCredentials) *Service {
	return &Service{
		creds: creds,
	}
}

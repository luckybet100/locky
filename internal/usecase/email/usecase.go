package email

import (
	"context"
	"github.com/Locky-Inc/locky/internal/domain/service/email"

	"github.com/google/uuid"

	"github.com/Locky-Inc/locky/internal/domain"
	"github.com/Locky-Inc/locky/internal/infrastructure/storage"
)

type Storage interface {
	ListEmails(ctx context.Context, filter storage.ListEmailsFilter) ([]*domain.Email, error)
	UpdateEmailStatus(ctx context.Context, id uuid.UUID, status domain.EmailSentStatus) error
}

type Adapter interface {
	SendEmail(ctx context.Context, creds *domain.EmailProviderCredentials, email *domain.Email) error
}

type SendWorkerConfig struct {
	WorkersCount int
	BatchesCount int64
	BatchSize    int64
}

type Usecase struct {
	st       Storage
	adapter  Adapter
	emailSvc email.Service
	SendWorkerConfig
}

func New(st Storage, adapter Adapter, emailSvc email.Service, cfg SendWorkerConfig) *Usecase {
	return &Usecase{st: st, SendWorkerConfig: cfg, emailSvc: emailSvc, adapter: adapter}
}

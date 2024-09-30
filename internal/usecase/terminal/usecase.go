package terminal

import (
	"context"
	"github.com/Locky-Inc/locky/internal/domain/service/terminal"

	"github.com/google/uuid"

	"github.com/Locky-Inc/locky/internal/domain"
	"github.com/Locky-Inc/locky/internal/infrastructure/storage"
)

type Storage interface {
	GetSpace(ctx context.Context, spaceID uuid.UUID) (*domain.Space, error)
	ListTerminals(ctx context.Context, filter storage.ListTerminalsFilter) ([]*domain.Terminal, error)
	CreateTerminal(ctx context.Context, terminal *domain.Terminal) error
}

type Usecase struct {
	st          Storage
	terminalSvc terminal.Service
}

func New(st Storage, terminalSvc terminal.Service) *Usecase {
	return &Usecase{
		terminalSvc: terminalSvc,
		st:          st,
	}
}

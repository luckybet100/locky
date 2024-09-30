package terminal

import (
	"github.com/google/uuid"
	
	"github.com/Locky-Inc/locky/internal/domain"
)

type CreateIn struct {
	SpaceID      uuid.UUID
	DisplayName  string
	SerialNumber string
	Model        domain.TerminalModel
}

func (s *Service) Create(in CreateIn) (*domain.Terminal, error) {
	terminal := &domain.Terminal{
		ID:          uuid.New(),
		DisplayName: in.DisplayName,
		TerminalInfo: domain.TerminalInfo{
			Model:        in.Model,
			SerialNumber: in.SerialNumber,
		},
	}
	if err := terminal.IsValid(); err != nil {
		return nil, err
	}
	return terminal, nil
}

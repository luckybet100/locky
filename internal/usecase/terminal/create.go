package terminal

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/Locky-Inc/locky/internal/domain"
	"github.com/Locky-Inc/locky/internal/domain/service/terminal"
	"github.com/Locky-Inc/locky/internal/infrastructure/storage"
)

type CreateIn struct {
	SpaceID      uuid.UUID
	DisplayName  string
	SerialNumber string
	Model        domain.TerminalModel
}

func (uc *Usecase) Create(ctx context.Context, in CreateIn) error {

	// создаем модель терминала и валидирует ввод
	term, err := uc.terminalSvc.Create(terminal.CreateIn(in))
	if err != nil {
		return errors.Join(err, ErrInvalidInput)
	}

	// проверяем наличие спейса 
	if _, getSpaceErr := uc.st.GetSpace(ctx, in.SpaceID); getSpaceErr != nil {
		if errors.Is(getSpaceErr, storage.ErrEntityNotFound) {
			return ErrSpaceNotFound
		}
		return fmt.Errorf("failed to get space: %w", getSpaceErr)
	}

	// добавляем терминал
	if createTerminalErr := uc.st.CreateTerminal(ctx, term); createTerminalErr != nil {
		if errors.Is(createTerminalErr, storage.ErrEntityExists) {
			return ErrTerminalAlreadyExists
		}
		return fmt.Errorf("failed to upsert terminal: %w", createTerminalErr)
	}

	return nil
}

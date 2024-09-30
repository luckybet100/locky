package terminal

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/Locky-Inc/locky/internal/domain"
	"github.com/Locky-Inc/locky/internal/infrastructure/storage"
)

type ListIn struct {
	SpaceID     uuid.UUID
	LastPageKey *uuid.UUID
	Limit       int64
}

type ListOut struct {
	Terminals   []*domain.Terminal
	NextPageKey *uuid.UUID
}

func (uc *Usecase) ListBySpace(ctx context.Context, in ListIn) (*ListOut, error) {

	// проверяем наличие спейса
	if _, err := uc.st.GetSpace(ctx, in.SpaceID); err != nil {
		if errors.Is(err, storage.ErrEntityNotFound) {
			return nil, ErrSpaceNotFound
		}
		return nil, fmt.Errorf("failed to get space: %w", err)
	}

	// получаем терминалы по спейсу
	terminals, err := uc.st.ListTerminals(ctx, storage.ListTerminalsFilter{
		SpaceID:     &in.SpaceID,
		LastPageKey: in.LastPageKey,
		Limit:       in.Limit,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get terminals by space id: %w", err)
	}

	var nextPageKey *uuid.UUID
	if len(terminals) != 0 {
		nextPageKey = &terminals[len(terminals)-1].ID
	}

	return &ListOut{
		Terminals:   terminals,
		NextPageKey: nextPageKey,
	}, nil
}

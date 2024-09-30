package space

import (
	"context"
	"errors"
	"fmt"

	"github.com/Locky-Inc/locky/internal/domain"
	"github.com/Locky-Inc/locky/internal/domain/service/space"
	"github.com/Locky-Inc/locky/internal/infrastructure/storage"
)

type CreateIn struct {
	DisplayName string
	Domain      string
	Admins      []domain.UserData
}

type CreateOut struct {
	Space  *domain.Space
	Admins []*domain.User
}

func (uc *Usecase) Create(ctx context.Context, in CreateIn) (*CreateOut, error) {
 
	data := uc.spaceSvc.Create(space.CreateIn(in))

	if err := uc.st.Transaction(ctx, func(ctx context.Context) error {

		if err := uc.st.CreateSpace(ctx, data.Space); err != nil {
			if errors.Is(err, storage.ErrEntityExists) {
				return ErrSpaceAlreadyExists
			}
			return fmt.Errorf("failed to create space: %w", err)
		}

		if err := uc.st.CreateUnit(ctx, data.StaffUnit); err != nil {
			return fmt.Errorf("failed to create staff unit for space: %w", err)
		}

		for _, admin := range data.Admins {
			if err := uc.st.CreateUser(ctx, admin); err != nil {
				return fmt.Errorf("failed to create space admin: %w", err)
			}
		}

		for _, email := range data.Emails {
			if err := uc.st.CreateEmail(ctx, email); err != nil {
				return fmt.Errorf("failed to create email: %w", err)
			}
		}

	}); err != nil {
		return nil, err
	}

	return &CreateOut{
		Space:  data.Space,
		Admins: data.Admins,
	}, nil
}

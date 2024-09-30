package space

import (
	"context"
	"github.com/Locky-Inc/locky/internal/domain"
	"github.com/Locky-Inc/locky/internal/domain/service/space"
)

type Storage interface {
	Transaction(ctx context.Context, txFunc func(ctx context.Context) error) error
	CreateSpace(ctx context.Context, space *domain.Space) error
	CreateUnit(ctx context.Context, unit *domain.Unit) error
	CreateUser(ctx context.Context, user *domain.User) error
	CreateEmail(ctx context.Context, email *domain.Email) error
}

type Usecase struct {
	st       Storage
	spaceSvc *space.Service
}

func New(st Storage, spaceSvc *space.Service) *Usecase {
	return &Usecase{
		st:       st,
		spaceSvc: spaceSvc,
	}
}

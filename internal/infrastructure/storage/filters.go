package storage

import (
	"context"
	"github.com/Locky-Inc/locky/internal/domain"
	"github.com/google/uuid"
)
type ListTerminalsFilter struct {
	SpaceID     *uuid.UUID
	LastPageKey *uuid.UUID
	Limit       int64
}

type ListEmailsFilter struct {
	Status      *domain.EmailSentStatus
	LastPageKey *uuid.UUID
	Limit       int64
}

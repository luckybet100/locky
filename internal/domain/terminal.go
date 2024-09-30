package domain

import (
	"errors"
	"github.com/google/uuid"
)

type TerminalModel string

const (
	VF618 = "vf618"
)

func (t TerminalModel) IsValid() error {
	switch t {
	case VF618:
		return nil
	default:
		return errors.New("unknown terminal model")
	}
}

/*
- уникальность displayName гарантируется в рамках одного спейса
- гарантируется глобальная уникальность serial number
*/
type Terminal struct {
	ID          uuid.UUID
	SpaceID     uuid.UUID
	DisplayName string
	TerminalInfo
}

func (t *Terminal) IsValid() error {
	// TODO добавить валиадцию остальных полей
	return t.TerminalInfo.IsValid()
}

type TerminalInfo struct {
	Model        TerminalModel
	SerialNumber string
}

func (t *TerminalInfo) IsValid() error {
	// TODO добавить валиадцию serial number
	return t.Model.IsValid()
}

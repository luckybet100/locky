package domain

import "github.com/google/uuid"

/*
под каждого клиента заводится отдельное пространство в рамках которого существуют пользователи
гарантируется уникальность display name
*/
type Space struct {
	ID          uuid.UUID
	StaffUnitID uuid.UUID
	DisplayName string
	Domain      string
}

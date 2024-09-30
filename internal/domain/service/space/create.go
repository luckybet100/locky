package space

import (
	"github.com/google/uuid"

	"github.com/Locky-Inc/locky/internal/domain"
)

type CreateIn struct {
	DisplayName string
	Domain      string
	Admins      []domain.UserData
}

type CreateOut struct {
	Space     *domain.Space
	StaffUnit *domain.Unit
	Admins    []*domain.User
	Emails    []*domain.Email
}

/*
Создаем новый пустой спейс и в нем заводим системный unit staff (без display name)
Добавляем админов в новый спейс и рассылаем им письма на создание пароля
*/
func (s *Service) Create(in CreateIn) CreateOut {
	staffUnitID := uuid.New()

	space := &domain.Space{
		ID:          uuid.New(),
		Domain:      in.Domain,
		DisplayName: in.DisplayName,
		StaffUnitID: staffUnitID,
	}

	admins := make([]*domain.User, 0, len(in.Admins))
	emails := make([]*domain.Email, 0, len(in.Admins))

	for _, admin := range in.Admins {
		user := &domain.User{
			ID:       uuid.New(),
			SpaceID:  space.ID,
			UnitID:   staffUnitID,
			UserData: admin,
		}
		admins = append(admins, user)
		emails = append(emails, domain.NewSetPasswordEmail(space, user))
	}

	return CreateOut{
		Space: space,
		StaffUnit: &domain.Unit{
			ID:      staffUnitID,
			SpaceID: space.ID,
		},
	}
}

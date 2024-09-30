package domain

import "github.com/google/uuid"

/*
пользователь спейса. гарантируется уникальность поля email (оно является обязательным, так как используется для входа)
*/

type User struct {
	ID      uuid.UUID
	SpaceID uuid.UUID
	UnitID  uuid.UUID
	UserData
}

type UserData struct {
	FirstName string
	LastName  string
	Email     string
}

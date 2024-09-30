package domain

import "github.com/google/uuid"

/*
организация, существующая в рамках спейса
display name -- ее отображаемое имя. уникальность display name гарантирована, если оно присуствует
если его нет, то unit не отображается в списке организаций и является системным (в данный момент есть только unit staff)
возможно в будущем добавим отдельный locky staff для админитсрации локи с отдельным акком для быстрого доступа к спейсу нашего персонала
*/
type Unit struct {
	ID          uuid.UUID
	SpaceID     uuid.UUID
	DisplayName *string
}

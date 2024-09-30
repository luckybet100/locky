package domain

import "github.com/google/uuid"

// Моделька сообщения, отправляемого на email

type EmailSentStatus string

const (
	EmailSentStatusPending      EmailSentStatus = "pending"
	EmailSentStatusFailedToSend EmailSentStatus = "failed"
	EmailSentStatusSent         EmailSentStatus = "sent"
)

type Email struct {
	ID      uuid.UUID
	To      string
	Subject string
	Body    string
	Status  EmailSentStatus
}

func NewSetPasswordEmail(space *Space, user *User) *Email {
	return &Email{
		ID:      uuid.New(),
		To:      user.Email,
		Subject: "Complete registration at locky.ai",
		Status:  EmailSentStatusPending,
	}
}

package domain

type EmailProviderType string

const (
	EmailProviderTypeSmtp EmailProviderType = "smtp"
	// TODO позднее реализуем фейкового провайдера для дебага на стейдже и e2e тестов
	EmailProviderTypeFake EmailProviderType = "fake" 
)

type EmailProviderCredentials struct {
	Type EmailProviderType
	SMTP *SmtpCredentials
}

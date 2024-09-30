package email

import (
	"context"

	"gopkg.in/gomail.v2"

	"github.com/Locky-Inc/locky/internal/domain"
)

func (c *Client) SendEmail(_ context.Context, creds *domain.SmtpCredentials, email *domain.Email) error {
	// TODO добавить логирование запросов и ошибок

	// получаем клиента до SMTP сервера
	dialer := c.getDialer(creds)

	// отправляем сообщение
	return dialer.DialAndSend(c.convertEmail(email, creds.Sender))
}

func (c *Client) convertEmail(email *domain.Email, sender string) *gomail.Message {
	msg := gomail.NewMessage()
	msg.SetHeader("From", sender)
	msg.SetHeader("To", email.To)
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody("text/plain", email.Body)
	return msg
}

func (c *Client) getDialer(creds *domain.SmtpCredentials) *gomail.Dialer {
	c.mu.Lock()
	defer c.mu.Unlock()

	// проверяем, если объект уже создан -- возвращаем его
	if dialer, ok := c.dialers[creds]; ok {
		return dialer
	}

	// создаем dialer для кредов, если ранее он не был создан
	dialer := gomail.NewDialer(creds.Host, creds.Port, creds.Login, creds.Password)
	c.dialers[creds] = dialer

	return dialer
}

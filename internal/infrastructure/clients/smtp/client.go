package email

import (
	"gopkg.in/gomail.v2"
	"sync"

	"github.com/Locky-Inc/locky/internal/domain"
)

type Client struct {
	dialers map[*domain.SmtpCredentials]*gomail.Dialer
	mu      sync.Mutex
}

func New() *Client {
	return &Client{
		dialers: make(map[*domain.SmtpCredentials]*gomail.Dialer),
	}
}

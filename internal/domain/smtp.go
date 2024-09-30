package domain

type SmtpCredentials struct {
	Host     string
	Port     int
	Login    string
	Password string
	Sender   string
}

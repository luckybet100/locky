package email

import "errors"

var ErrNoWorkingProvider = errors.New("failed to find working provider to send email")

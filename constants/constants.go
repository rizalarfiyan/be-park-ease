package constants

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

var (
	FiberBodyLimit   = 5 * 1024 * 1024 // 5 MB
	DefaultPageLimit = 10
)

var (
	ValidationUsername = validation.Match(regexp.MustCompile(`^[a-zA-Z0-9._][a-zA-Z0-9._]{3,18}[a-zA-Z0-9._]$`)).Error("the invalid format, must be 5-20 characters, only alphanumeric, dot, and underscore")
	ValidationPassword = validation.Match(regexp.MustCompile(`^[A-Za-z\d@$!%*#?&]*[A-Za-z][A-Za-z\d@$!%*#?&]*\d[A-Za-z\d@$!%*#?&]*[@$!%*#?&][A-Za-z\d@$!%*#?&]*$`)).Error("the invalid format, must be 8 characters, at least 1 letter, 1 number and 1 special character")
)

package request

import (
	"be-park-ease/constants"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AuthLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req AuthLoginRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username, validation.Required, constants.ValidationUsername.Error("Invalid username or password")),
		validation.Field(&req.Password, validation.Required, constants.ValidationPassword.Error("Invalid username or password")),
	)
}

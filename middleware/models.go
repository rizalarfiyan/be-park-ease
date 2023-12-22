package middleware

import (
	"be-park-ease/internal/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type AuthUserData struct {
	ID       int32          `json:"id"`
	Name     string         `json:"name"`
	Username string         `json:"username"`
	Role     sql.UserRole   `json:"role"`
	Status   sql.UserStatus `json:"status"`
}

func (data *AuthUserData) Get(ctx *fiber.Ctx) error {
	user, isvalid := ctx.Locals("user").(AuthUserData)
	if !isvalid {
		return fmt.Errorf("could not extract user from context")
	}

	*data = user
	return nil
}

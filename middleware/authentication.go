package middleware

import (
	"be-park-ease/internal/sql"
	"be-park-ease/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (m *middleware) Auth(isList bool) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")
		if utils.IsEmpty(token) {
			m.exception.IsUnauthorize(MsgAuthExpired, isList)
		}

		user, err := m.repo.GetUserByToken(ctx.Context(), token)
		m.exception.PanicIfErrorWithoutNoSqlResult(err, false)
		m.exception.IsBadRequest(user, MsgAuthExpired, false)

		isBanned := user.Status != sql.UserStatusBanned
		m.exception.IsUnprocessableEntity(isBanned, MsgAuthDisable, false)

		if user.ExpiredAt.Valid && user.ExpiredAt.Time.Before(time.Now()) {
			m.exception.IsUnauthorize(MsgAuthExpired, isList)
		}

		payload := AuthUserData{
			ID:       user.ID,
			Name:     user.Name,
			Username: user.Username,
			Role:     user.Role,
			Status:   user.Status,
		}

		ctx.Locals("user", payload)

		return ctx.Next()
	}
}

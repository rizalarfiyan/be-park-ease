package middleware

import (
	"be-park-ease/internal/sql"

	"github.com/gofiber/fiber/v2"
)

func (m *middleware) baseRoles(roles []sql.UserRole, isList bool) fiber.Handler {
	mapRoles := make(map[sql.UserRole]bool)

	for _, role := range roles {
		if _, ok := mapRoles[role]; !ok {
			mapRoles[role] = true
		}
	}

	return func(ctx *fiber.Ctx) error {
		user := AuthUserData{}
		err := user.Get(ctx)
		if err != nil {
			m.exception.IsUnauthorize(MsgAuthNotImplement, isList)
		}

		if _, ok := mapRoles[user.Role]; !ok {
			m.exception.IsForbidden(MsgAuthNotImplement, isList)
		}

		return ctx.Next()
	}
}

func (m *middleware) Role(role sql.UserRole, isList bool) fiber.Handler {
	return m.baseRoles([]sql.UserRole{role}, isList)
}

func (m *middleware) Roles(roles []sql.UserRole, isList bool) fiber.Handler {
	return m.baseRoles(roles, isList)
}

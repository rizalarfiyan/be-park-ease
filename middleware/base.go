package middleware

import (
	"be-park-ease/config"
	"be-park-ease/exception"
	"be-park-ease/internal/repository"
	"be-park-ease/internal/sql"

	"github.com/gofiber/fiber/v2"
)

type Middleware interface {
	Auth(isList bool) fiber.Handler
	Role(role sql.UserRole, isList bool) fiber.Handler
	Roles(roles []sql.UserRole, isList bool) fiber.Handler
}

type middleware struct {
	repo      repository.AuthRepository
	conf      *config.Config
	exception exception.Exception
}

func NewMiddleware(repo repository.AuthRepository) Middleware {
	return &middleware{
		repo:      repo,
		conf:      config.Get(),
		exception: exception.NewException("middleware"),
	}
}

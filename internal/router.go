package internal

import (
	"be-park-ease/internal/handler"

	"github.com/gofiber/fiber/v2"
)

type Router interface {
	BaseRoute(handler handler.BaseHandler)
}

type router struct {
	app *fiber.App
}

func NewRouter(app *fiber.App) Router {
	return &router{
		app: app,
	}
}

func (r *router) BaseRoute(handler handler.BaseHandler) {
	r.app.Get("", handler.Home)
}

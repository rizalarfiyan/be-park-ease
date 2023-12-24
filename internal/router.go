package internal

import (
	"be-park-ease/internal/handler"
	"be-park-ease/middleware"

	"github.com/gofiber/fiber/v2"
)

type Router interface {
	BaseRoute(handler handler.BaseHandler)
	AuthRoute(handler handler.AuthHandler)
	HistoryRoute(handler handler.HistoryHandler)
}

type router struct {
	app *fiber.App
	mid middleware.Middleware
}

func NewRouter(app *fiber.App, middleware middleware.Middleware) Router {
	return &router{
		app: app,
		mid: middleware,
	}
}

func (r *router) BaseRoute(handler handler.BaseHandler) {
	r.app.Get("", handler.Home)
}

func (r *router) AuthRoute(handler handler.AuthHandler) {
	auth := r.app.Group("auth")
	auth.Get("me", r.mid.Auth(false), handler.Me)
	auth.Post("login", handler.Login)
}

func (r *router) HistoryRoute(handler handler.HistoryHandler) {
	history := r.app.Group("history")
	entry := history.Group("entry")
	entry.Get("", r.mid.Auth(false), handler.AllHistory)
}

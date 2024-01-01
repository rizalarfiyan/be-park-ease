package internal

import (
	"be-park-ease/internal/handler"
	"be-park-ease/internal/sql"
	"be-park-ease/middleware"

	"github.com/gofiber/fiber/v2"
)

type Router interface {
	BaseRoute(handler handler.BaseHandler)
	AuthRoute(handler handler.AuthHandler)
	HistoryRoute(handler handler.HistoryHandler)
	UserRoute(handler handler.UserHandler)
	SettingRoute(handler handler.SettingHandler)
	VehicleTypeRoute(handler handler.VehicleTypeHandler)
	LocationRoute(handler handler.LocationHandler)
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
	history.Get("", r.mid.Auth(false), handler.AllHistory)
	history.Post("entry", r.mid.Auth(false), handler.CreateEntryHistory)
	history.Post("calculate", r.mid.Auth(false), handler.CalculatePriceHistory)
}

func (r *router) UserRoute(handler handler.UserHandler) {
	user := r.app.Group("user")
	user.Get("", r.mid.Auth(false), r.mid.Role(sql.UserRoleAdmin, false), handler.AllUser)
	user.Post("", r.mid.Auth(false), r.mid.Role(sql.UserRoleAdmin, false), handler.CreateUser)
	user.Post("change-password", r.mid.Auth(false), handler.ChangePassword)
	user.Get(":id", r.mid.Auth(false), r.mid.Role(sql.UserRoleAdmin, false), handler.UserById)
	user.Put(":id", r.mid.Auth(false), r.mid.Role(sql.UserRoleAdmin, false), handler.UpdateUser)
}

func (r *router) SettingRoute(handler handler.SettingHandler) {
	setting := r.app.Group("setting")
	setting.Get("", r.mid.Auth(false), r.mid.Role(sql.UserRoleAdmin, false), handler.GetAllSetting)
	setting.Post("", r.mid.Auth(false), r.mid.Role(sql.UserRoleAdmin, false), handler.CreateOrUpdateSetting)
}

func (r *router) VehicleTypeRoute(handler handler.VehicleTypeHandler) {
	vehicleType := r.app.Group("vehicle_type")
	vehicleType.Get("", r.mid.Auth(true), r.mid.Role(sql.UserRoleAdmin, true), handler.AllVehicleType)
	vehicleType.Post("", r.mid.Auth(false), r.mid.Role(sql.UserRoleAdmin, false), handler.CreateVehicleType)
	vehicleType.Get(":code", r.mid.Auth(false), r.mid.Role(sql.UserRoleAdmin, false), handler.VehicleTypeByCode)
	vehicleType.Put(":code", r.mid.Auth(false), r.mid.Role(sql.UserRoleAdmin, false), handler.UpdateVehicleType)
	vehicleType.Delete(":code", r.mid.Auth(false), r.mid.Role(sql.UserRoleAdmin, false), handler.DeleteVehicleType)
}

func (r *router) LocationRoute(handler handler.LocationHandler) {
	location := r.app.Group("location")
	location.Get("", r.mid.Auth(true), r.mid.Role(sql.UserRoleAdmin, true), handler.AllLocation)
	location.Post("", r.mid.Auth(false), r.mid.Role(sql.UserRoleAdmin, false), handler.CreateLocation)
	location.Get(":code", r.mid.Auth(false), r.mid.Role(sql.UserRoleAdmin, false), handler.LocationByCode)
	location.Put(":code", r.mid.Auth(false), r.mid.Role(sql.UserRoleAdmin, false), handler.UpdateLocation)
	location.Delete(":code", r.mid.Auth(false), r.mid.Role(sql.UserRoleAdmin, false), handler.DeleteLocation)
}

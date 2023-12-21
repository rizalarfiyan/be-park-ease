package config

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
)

func CorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     conf.Cors.AllowOrigins,
		AllowMethods:     conf.Cors.AllowMethods,
		AllowHeaders:     conf.Cors.AllowHeaders,
		AllowCredentials: conf.Cors.AllowCredentials,
		ExposeHeaders:    conf.Cors.ExposeHeaders,
	}
}

func FiberZerolog(rawLogs *zerolog.Logger, logs zerolog.Logger) fiberzerolog.Config {
	fields := []string{
		fiberzerolog.FieldIP,
		fiberzerolog.FieldMethod,
		fiberzerolog.FieldPath,
		fiberzerolog.FieldURL,
		fiberzerolog.FieldMethod,
		fiberzerolog.FieldPath,
		fiberzerolog.FieldLatency,
		fiberzerolog.FieldStatus,
		fiberzerolog.FieldBody,
		fiberzerolog.FieldError,
		fiberzerolog.FieldRequestID,
	}

	return fiberzerolog.Config{
		Logger: &logs,
		Fields: fields,
		SkipBody: func(ctx *fiber.Ctx) bool {
			return strings.Contains(string(ctx.Request().Header.ContentType()), "multipart/form-data")
		},
		GetLogger: func(c *fiber.Ctx) zerolog.Logger {
			call, ok := c.Locals("caller").(string)
			if !ok {
				return logs
			}

			return rawLogs.With().Str("caller", call).Logger()
		},
	}
}

func FiberRecover() recover.Config {
	return recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			_, file, line, ok := runtime.Caller(4)
			if !ok {
				return
			}
			c.Locals("caller", fmt.Sprintf("%s:%d", file, line))
		},
	}
}

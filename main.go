package main

import (
	_ "be-park-ease/docs"

	"be-park-ease/config"
	"be-park-ease/internal"
	"be-park-ease/internal/handler"
	"be-park-ease/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog"
)

func init() {
	config.Init()
	conf := config.Get()
	logger.Init(conf)
	logger.UpdateLogLevel(conf.Logger.Level)
}

func main() {
	conf := config.Get()
	rawLogs := logger.GetWithoutCaller("main-api")
	logs := rawLogs.With().Caller().Logger()

	app := fiber.New()
	app.Use(fiberzerolog.New(config.FiberZerolog(rawLogs, logs)))
	app.Use(recover.New(config.FiberRecover()))
	app.Use(requestid.New())
	app.Use(cors.New(config.CorsConfig()))
	app.Use(compress.New())
	app.Use(helmet.New())

	app.Get("/swagger/*", basicauth.New(basicauth.Config{
		Users: map[string]string{
			conf.Swagger.Username: conf.Swagger.Password,
		},
	}), swagger.New(swagger.Config{
		URL:          "/swagger/doc.json",
		DeepLinking:  true,
		DocExpansion: "list",
	}))

	route := internal.NewRouter(app)

	// handler
	baseHandler := handler.NewBaseHandler()

	// router
	route.BaseRoute(baseHandler)

	baseUrl := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	server := &http.Server{
		Addr: baseUrl,
	}

	go func() {
		err := app.Listen(baseUrl)
		if err != nil {
			logs.Fatal().Err(err).Msg("Error app serve")
		}
	}()

	handleShutdown(server, logs)
}

func handleShutdown(server *http.Server, logs zerolog.Logger) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logs.Warn().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	if err = server.Shutdown(ctx); err != nil {
		logs.Err(err).Msg("Server forced to shutdown")
	}

	logs.Info().Msg("Server exiting")
}

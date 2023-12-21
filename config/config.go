package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

var conf *Config

func Init() {
	cc := NewConfigConvert()
	appEnv := cc.AsString("APP_ENV", "development")

	err := godotenv.Load(".env")
	if err != nil && appEnv != "production" {
		log.Fatal(".env is not loaded properly. Err: ", err)
	}

	conf = new(Config)
	conf.Host = cc.AsString("HOST", "")
	conf.Port = cc.AsInt("PORT", 8080)
	conf.Name = cc.AsString("APP_NAME", "Park Ease")

	conf.Logger.Level = cc.AsZerologLevel("LOG_LEVEL", zerolog.InfoLevel)
	conf.Logger.Path = cc.AsString("LOG_PATH", "./log/logger.log")
	conf.Logger.IsCompressed = cc.AsBool("LOG_COMPRESSED", true)
	conf.Logger.IsDailyRotate = cc.AsBool("LOG_DAILY_ROTATE", true)
	conf.Logger.IsLogRotator = cc.AsBool("LOG_ROTATOR", true)
	conf.Logger.SleepDuration = cc.AsTimeDuration("LOG_SLEEP_DURATION", 5*time.Second)

	conf.Cors.AllowOrigins = cc.AsString("ALLOW_ORIGINS", "*")
	conf.Cors.AllowMethods = cc.AsString("ALLOW_METHODS", "GET,POST,PUT,DELETE,OPTIONS")
	conf.Cors.AllowHeaders = cc.AsString("ALLOW_HEADERS", "Origin,Content-Type,Accept,Authorization")
	conf.Cors.AllowCredentials = cc.AsBool("ALLOW_CREDENTIALS", false)
	conf.Cors.ExposeHeaders = cc.AsString("EXPOSE_HEADERS", "Content-Length,Content-Type,Authorization")

	conf.DB.Name = cc.AsString("DB_NAME", "app")
	conf.DB.Host = cc.AsString("DB_HOST", "localhost")
	conf.DB.Port = cc.AsInt("DB_PORT", 3306)
	conf.DB.User = cc.AsString("DB_USER", "root")
	conf.DB.Password = cc.AsString("DB_PASSWORD", "password")
	conf.DB.ConnectionIdle = cc.AsTimeDuration("DB_CONNECTION_IDLE", 1*time.Minute)
	conf.DB.ConnectionLifetime = cc.AsTimeDuration("DB_CONNECTION_LIFETIME", 5*time.Minute)
	conf.DB.MaxIdle = cc.AsInt("DB_MAX_IDLE", 20)
	conf.DB.MaxOpen = cc.AsInt("DB_MAX_OPEN", 50)

	conf.Swagger.Username = cc.AsString("SWAGGER_USERNAME", "admin")
	conf.Swagger.Password = cc.AsString("SWAGGER_PASSWORD", "password")

	conf.Auth.TokenSalt = cc.AsString("TOKEN_SALT", "")
}

func Get() *Config {
	return conf
}

package config

import (
	"time"

	"github.com/rs/zerolog"
)

type Config struct {
	Env     string
	Name    string
	Port    int
	Host    string
	Logger  LoggerConfigs
	Cors    CorsConfigs
	Swagger SwaggerConfigs
	DB      DBConfigs
	Auth    AuthConfigs
}

type LoggerConfigs struct {
	Level         zerolog.Level
	Path          string
	IsCompressed  bool
	IsDailyRotate bool
	SleepDuration time.Duration
}

type CorsConfigs struct {
	AllowOrigins     string
	AllowMethods     string
	AllowHeaders     string
	AllowCredentials bool
	ExposeHeaders    string
}

type DBConfigs struct {
	Host               string
	Port               int
	User               string
	Password           string
	Name               string
	ConnectionIdle     time.Duration
	ConnectionLifetime time.Duration
	MaxIdle            int
	MaxOpen            int
}

type SwaggerConfigs struct {
	Username string
	Password string
}

type AuthConfigs struct {
	TokenSalt string
}

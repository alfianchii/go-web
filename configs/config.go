package configs

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
)

var (
	ErrFailedLoadENVFile = errors.New("failed loading .env file")
	ErrFailedReadENVFile = errors.New("failed reading .env file")
)

type Config struct {
	AppName string
	AppURL string
	AppPort string

	DBDriver string
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string
	
	JWTSecret string
}

var (
	ExecTimeoutDuration = 10*time.Second
	TokenDuration = 1*time.Hour
)

func InitENV() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(ErrFailedLoadENVFile)
	}

	return &Config{
		AppName: GetENV("APP_NAME"),
		AppURL: GetENV("APP_URL"),
		AppPort: GetENV("APP_PORT"),
		DBDriver: GetENV("DB_DRIVER"),
		DBHost: GetENV("DB_HOST"),
		DBPort: GetENV("DB_PORT"),
		DBName: GetENV("DB_DATABASE"),
		DBUser: GetENV("DB_USERNAME"),
		DBPass: GetENV("DB_PASSWORD"),
		JWTSecret: GetENV("JWT_SECRET"),
	}
}

func GetENV(key string) string {
	dotEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal(ErrFailedReadENVFile)
	}

	return dotEnv[key]
}

func GetAppAddress(cfg *Config) string {
	return fmt.Sprintf("%s:%s", cfg.AppURL, cfg.AppPort)
}

func CtxTime() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(CtxBg(), ExecTimeoutDuration)
	
	return ctx, cancel
}

func CtxBg() context.Context {
	return context.Background()
}
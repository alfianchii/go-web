package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
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
	Address = fmt.Sprintf("%s:%s", GetENV("APP_URL"), GetENV("APP_PORT"))
	ExecTimeoutDuration = 10*time.Second
	TokenDuration = 1*time.Hour
)

func GetENV(key string) string {
	dotEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error reading .env file")
	}

	return dotEnv[key]
}

func CtxTime() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(CtxBg(), ExecTimeoutDuration)
	
	return ctx, cancel
}

func CtxBg() context.Context {
	return context.Background()
}
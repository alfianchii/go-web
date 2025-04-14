package database

import (
	"fmt"
	"go-web/configs"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func InitDB(config *configs.Config) *DB {
	dsn := dbDSN(config)

	pool, err := pgxpool.New(configs.CtxBg(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	ctx, cancel := configs.CtxTime()
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	fmt.Println("Successfully connected to the PostgreSQL database!")

	return &DB{Pool: pool}
}

func dbDSN (config *configs.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBName)
}
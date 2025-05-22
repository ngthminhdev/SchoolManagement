package db

import (
	"GolangBackend/config"
	"GolangBackend/internal/global"
	"GolangBackend/helper"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDatabase() {
	var host string = config.GetEnv("DB_HOST", "localhost")
	var port string = config.GetEnv("DB_PORT", "5432")
	var user string = config.GetEnv("DB_USER", "admin")
	var password string = config.GetEnv("DB_PASSWORD", "password")
	var dbname string = config.GetEnv("DB_NAME", "school_management")

	var dns string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s pool_max_conns=10 sslmode=disable", host, port, user, password, dbname)

	ctx, cancelConnect := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelConnect()

	config, err := pgxpool.ParseConfig(dns)

	if err != nil {
		log.Fatalf("Error when connect database: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		log.Fatalf("Error when create connection pool: %v", err)
	}

	global.DB = pool
	helper.LogInfo("Connected to database")

	RunMigrations()
}

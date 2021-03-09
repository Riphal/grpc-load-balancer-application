package grpc_load_balancer_application

import (
	"log"

	"github.com/Riphal/grpc-load-balancer-application/common/config"
	"github.com/Riphal/grpc-load-balancer-application/common/storage/postgres"
	"github.com/Riphal/grpc-load-balancer-application/common/storage/redis"
)

type App struct {
	DB 		*postgres.DB
	Redis 	*redis.Storage
}

func NewApp() *App {
	err := config.LoadEnvFile("")
	if err != nil {
		log.Println("Failed to load env file", err)
	}

	return &App{
		DB: 	mustInitDB(),
		Redis: 	mustInitRedis(),
	}
}

func mustInitDB() *postgres.DB {
	db, err := postgres.New(config.GetEnv("DATABASE_URL", "postgres://postgres:example@localhost:5155/postgres?sslmode=disable"))
	if err.IsNotNil() {
		log.Fatalln("failed to initialize postgres", err)
	}

	log.Println("[init] postgres initialized")

	return db
}

func mustInitRedis() *redis.Storage {
	redisConfig := &redis.Config{
		ConnectionURL: config.GetEnv("REGISTRY_URL", "redis://localhost:6300"),
	}

	redisStorage, err := redis.New(redisConfig)
	if err.IsNotNil() {
		log.Fatalln("failed to initialize redis", err)
	}

	log.Println("[init] redis initialized")

	return redisStorage
}

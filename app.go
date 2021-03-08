package grpc_load_balancer_application

import (
	"log"

	"github.com/Riphal/grpc-load-balancer-application/common/config"
	"github.com/Riphal/grpc-load-balancer-application/common/storage/postgres"
	"github.com/Riphal/grpc-load-balancer-application/common/storage/redis"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/server"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/server/router"
)

type App struct {
	Server	*server.Server

	DB 		*postgres.DB
	Redis 	*redis.Storage
}

func NewApp() *App {
	return &App{
		Server:	initServer(),
		DB: 	initDB(),
		Redis: 	initRedis(),
	}
}

func initServer() *server.Server {
	r := router.New()
	port := ":" + config.GetEnv("PORT", "3001")

	serverConfig := &server.Config{
		Address: port,
		Router:  r,
	}

	return server.New(serverConfig)
}

func initDB() *postgres.DB {
	db, err := postgres.New(config.GetEnv("DATABASE_URL", "postgres://postgres:example@localhost:5155/postgres"))
	if err.IsNotNil() {
		log.Fatalln("failed to initialize postgres", err)
	}

	log.Println("[init] postgres initialized")

	return db
}

func initRedis() *redis.Storage {
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

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Muhammadjon226/user_service/api"
	"github.com/Muhammadjon226/user_service/config"
	"github.com/Muhammadjon226/user_service/events"
	"github.com/Muhammadjon226/user_service/pkg/logger"
	"github.com/Muhammadjon226/user_service/service"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	if info, err := os.Stat(".env"); !os.IsNotExist(err) {
		if !info.IsDir() {
			godotenv.Load(".env")
		}
	}
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "user_service")
	defer logger.Cleanup(log)

	log.Info("main: pgxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase)

	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		log.Error("postgres connect error",
			logger.Error(err))
		return
	}
	defer connDb.Close()

	userService := service.NewUserService(connDb, log, &cfg )
	server := api.New(api.Config{
		Logger:      log,
		Config:      cfg,
		UserService: userService,
	})

	pubsubServer, err := events.New(cfg, log, connDb)
	if err != nil {
		fmt.Println("errrror in pubsub server")
		log.Fatal("Error on event server", logger.Error(err))
	}

	if cfg.Environment == "develop" {
		go func() {
			pubsubServer.Run(context.Background())
			log.Fatal("Event server has finished")

		}()
	}

	fmt.Println("port: ", cfg.RPCPort)
	if err := server.Run(cfg.RPCPort); err != nil {
		log.Fatal("error while running gin server", logger.Error(err))
	}
	log.Fatal("API server has finished")
}

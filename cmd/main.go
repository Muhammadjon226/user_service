package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/Muhammadjon226/user_service/config"
	"github.com/Muhammadjon226/user_service/events"
	pb "github.com/Muhammadjon226/user_service/genproto/user_service"
	"github.com/Muhammadjon226/user_service/pkg/logger"
	"github.com/Muhammadjon226/user_service/service"
	grpcclient "github.com/Muhammadjon226/user_service/service/grpcclient"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	client, err := grpcclient.New(cfg)
	if err != nil {
		log.Error("error while connecting other services")
		return
	}

	userService := service.NewUserService(connDb, log, client, &cfg)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()

	// err = sentry.Init(sentry.ClientOptions{
	// 	Dsn: cfg.SentryDNS,
	// })
	// if err != nil {
	// 	log.Fatal("Cannot initialize sentry.io", logger.Error(err))
	// }

	pb.RegisterUserServiceServer(s, userService)

	reflection.Register(s)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	pubsubServer, err := events.New(cfg, log, connDb, client)
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
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
	log.Fatal("API server has finished")
}

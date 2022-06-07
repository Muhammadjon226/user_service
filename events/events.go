package events


import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/Muhammadjon226/user_service/config"
	"github.com/Muhammadjon226/user_service/events/user_service/user"
	"github.com/Muhammadjon226/user_service/pkg/event"
	"github.com/Muhammadjon226/user_service/pkg/logger"
	grpcClient"github.com/Muhammadjon226/user_service/service/grpcclient"
)

// PubsubServer ...
type PubsubServer struct {
	cfg   config.Config
	log   logger.Logger
	db    *sqlx.DB
	kafka *event.Kafka
	newService grpcClient.IServiceManager
}

// New ...
func New(cfg config.Config, log logger.Logger, db *sqlx.DB, newSer grpcClient.IServiceManager) (*PubsubServer, error) {

	kafka, err := event.NewKafka(cfg, log)
	if err != nil {
		return nil, err
	}

	// kafka.AddPublisher("v1.websocket_service.response")
	return &PubsubServer{
		cfg:   cfg,
		log:   log,
		db:    db,
		kafka: kafka,
		newService:    newSer,
	}, nil
}

// Run ...
func (s *PubsubServer) Run(ctx context.Context) {

	postService := user.NewService(s.db, s.log, &s.cfg, s.kafka)
	postService.RegisterConsumers()
	
	s.kafka.RunConsumers(ctx)
}

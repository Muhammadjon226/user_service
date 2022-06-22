package user

import (
	"github.com/Muhammadjon226/user_service/config"
	"github.com/Muhammadjon226/user_service/storage"
	"github.com/jmoiron/sqlx"

	"github.com/Muhammadjon226/user_service/pkg/event"
	"github.com/Muhammadjon226/user_service/pkg/logger"
)

// Service ...
type Service struct {
	storage storage.IStorage
	logger  logger.Logger
	config  *config.Config
	kafka   *event.Kafka
}

//NewService ...
func NewService(db *sqlx.DB, log logger.Logger, config *config.Config, kafka *event.Kafka) *Service {
	return &Service{
		storage: storage.NewStoragePg(db),
		logger:  log,
		config:  config,
		kafka:   kafka,
	}
}

//RegisterConsumers ...
func (c *Service) RegisterConsumers() {
	adRoute := "v1.user_serivce.user."

	c.kafka.AddConsumer(
		adRoute+"created", 	 // consumer name
		"v1.user.created",   // topic
		adRoute+"created", 	 // group id
		c.Created,         	 // handlerFunction
	)
}

package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/Muhammadjon226/user_service/config"

	"github.com/Muhammadjon226/user_service/pkg/logger"
	grpcclient "github.com/Muhammadjon226/user_service/service/grpcclient"
	"github.com/Muhammadjon226/user_service/storage"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  logger.Logger
	client  grpcclient.IServiceManager
	config  *config.Config
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log logger.Logger, client grpcclient.IServiceManager, config *config.Config) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
		config:  config,
	}
}

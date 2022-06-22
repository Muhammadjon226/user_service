package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/Muhammadjon226/user_service/config"

	"github.com/Muhammadjon226/user_service/pkg/logger"
	"github.com/Muhammadjon226/user_service/storage"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  logger.Logger
	config  *config.Config
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log logger.Logger, config *config.Config) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		config:  config,
	}
}

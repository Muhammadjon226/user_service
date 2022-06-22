package v1

import (
	"github.com/Muhammadjon226/user_service/config"
	"github.com/Muhammadjon226/user_service/pkg/logger"
	"github.com/Muhammadjon226/user_service/service"
)

//HandlerV1 ...
type HandlerV1 struct {
	log   logger.Logger
	cfg   config.Config
	userService *service.UserService
}

// New ...
func New(log logger.Logger, cfg config.Config, userService *service.UserService) *HandlerV1 {

	return &HandlerV1{
		cfg:   cfg,
		log:   log,
		userService: userService,
	}
}

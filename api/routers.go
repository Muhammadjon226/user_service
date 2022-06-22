package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// _ "github.com/Muhammadjon226/user_service/api/docs" // swag
	v1 "github.com/Muhammadjon226/user_service/api/handlers/v1"
	"github.com/Muhammadjon226/user_service/config"
	"github.com/Muhammadjon226/user_service/pkg/logger"
	"github.com/Muhammadjon226/user_service/service"
)

type Config struct {
	Logger      logger.Logger
	Config      config.Config
	UserService *service.UserService
}

// New is a constructor for gin.Engine
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cfg Config) *gin.Engine {
	if cfg.Config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// this html has been copied to Dockerfile
	// r.LoadHTMLGlob("html/**/*")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	handlerV1 := v1.New(
		cfg.Logger,
		cfg.Config,
		cfg.UserService,
	)

	r.GET("/v1/user/list-users/", handlerV1.ListUsers)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}

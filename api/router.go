package api

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/FarrukhibnAkbar/bookshop-api/api/handlers/v1"
	"github.com/FarrukhibnAkbar/bookshop-api/config"
	"github.com/FarrukhibnAkbar/bookshop-api/pkg/logger"
	"github.com/FarrukhibnAkbar/bookshop-api/services"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

func New(option Option) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("v1")
	api.POST("/auhtors", handlerV1.CreateAuthor)

	return router
}

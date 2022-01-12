package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/FarrukhibnAkbar/bookshop-api/api/docs" // swag
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
	api.POST("/authors", handlerV1.CreateAuthor)
	api.GET("/authors/", handlerV1.GetAuthor)
	api.GET("/authors", handlerV1.ListAuthors)
	api.PUT("/authors/:id", handlerV1.UpdateAuthor)
	api.DELETE("/authors/:id", handlerV1.DeleteAuthor)

	api.POST("/books", handlerV1.CreateBook)
	api.GET("/books/{id}", handlerV1.GetBook)
	api.GET("/books", handlerV1.ListBooks)
	api.PUT("/books/{id}", handlerV1.UpdateBook)
	api.DELETE("/books/{id}", handlerV1.DeleteBook)

	api.POST("/categories", handlerV1.CreateCategory)
	api.GET("/categories/}", handlerV1.GetCategory)
	api.GET("/categories", handlerV1.ListCategories)
	api.PUT("/categories/:id", handlerV1.UpdateCategory)
	api.DELETE("/categories/:id", handlerV1.DeleteCategory)

	api.POST("/orders", handlerV1.CreateOrder)
	api.GET("/orders/", handlerV1.GetOrder)
	api.GET("/orders", handlerV1.ListOrders)
	api.PUT("/orders/:id", handlerV1.UpdateOrder)
	api.DELETE("/orders/:id", handlerV1.DeleteOrder)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

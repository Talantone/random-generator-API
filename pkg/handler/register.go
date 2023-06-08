package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "random-generator-API/docs"
)

func (h *Handler) RegisterHTTPEndpoints() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	generator := router.Group("/generator")
	{
		generator.POST("/generate", h.Generate)
		generator.GET("/result/:id", h.Result)
	}
	return router
}

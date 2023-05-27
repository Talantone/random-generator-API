package handler

import (
	"github.com/gin-gonic/gin"
	"random-generator-API/models"
	"random-generator-API/pkg"
)

func RegisterHTTPEndpoints[T models.RandomItemTypes](router *gin.RouterGroup, uc pkg.UseCase[T]) {
	h := NewHandler(uc)
	generator := router.Group("/generator")
	{
		generator.POST("/generate", h.Generate)
		generator.GET("/result", h.Result)
	}
}

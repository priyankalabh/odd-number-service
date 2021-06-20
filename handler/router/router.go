package router

import (
	"odd-number-service/handler"

	"github.com/gin-gonic/gin"
)

func AttachRouter(router *gin.Engine) {
	router.GET("/api/odd", handler.OddHandler)
}

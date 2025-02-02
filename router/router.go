package router

import (
	"examplegpt.com/controller"
	"github.com/gin-gonic/gin"
)

func Router(eng *gin.Engine) {
	eng.POST("/pergunta", controller.InputPergunta)
	eng.GET("/", controller.ServeIndex)
}

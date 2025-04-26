package routes

import (
	"crudgo/handler"

	"github.com/gin-gonic/gin"
)

// Função para configurar as rotas do usuário
func UserRoutes(router *gin.Engine) {

	router.GET("/user", handler.UserGet)
	router.POST("/user", handler.UserPost)
	router.GET("/user/:id", handler.UserGetById)
	router.DELETE("/user/:id", handler.UserDelete)

}

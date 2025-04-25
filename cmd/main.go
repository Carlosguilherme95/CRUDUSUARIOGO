package main

import (
	"crudgo/config"
	"crudgo/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar o banco de dados
	config.InitDB()
	router := gin.Default()
	routes.UserRoutes(router)

	log.Println("Servidor iniciado na porta 8080")

	// Verificando o erro ao iniciar o servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}

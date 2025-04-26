package handler

import (
	"crudgo/entities"
	"crudgo/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserGet(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Method not allowed"})
		return
	}
	c.JSON(http.StatusOK, users)
	log.Println("Usuários retornados com sucesso")
}
func UserGetById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	// Chama o serviço para obter o usuário pelo ID
	user, err := service.GetOneUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, user)
	log.Println("Usuário retornado com sucesso")
}
func UserDelete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	err = service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "usuário não encontrado"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "usuário deletado com sucesso"})
	log.Println("Usuário deletado com sucesso")
}

// Função para lidar com a requisição POST de criação de usuário
func UserPost(C *gin.Context) {
	var user entities.User
	if err := C.ShouldBindJSON(&user); err != nil {
		C.JSON(400, gin.H{"error": "dados inválidos"})
	}
	createdUser, err := service.CreateUser(user.FirstName, user.LastName, user.Email)
	if err != nil {
		C.JSON(500, gin.H{"error": "erro ao criar usuário"})
		return
	}
	C.JSON(200, gin.H{"message": "usuário criado com sucesso", "usuario": createdUser})

}

package service

import (
	"crudgo/config"
	"crudgo/entities"
	"errors"
	"log"
)

// Verifica se os campos são válidos e se o e-mail já está cadastrado
func checkUserAndEmail(firstName, lastName, email string) error {
	db := config.GetDB()

	if firstName == "" || lastName == "" || email == "" {
		return errors.New("todos os campos são obrigatórios")
	}

	var existingUser entities.User
	if err := db.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return errors.New("usuário com este email já existe")
	}

	return nil
}
func GetAllUsers() ([]entities.User, error) {
	var users []entities.User
	db := config.GetDB()
	result := db.Find(&users)
	if result.Error != nil {
		log.Println("Erro ao buscar usuários:", result.Error)
	}
	return users, result.Error
}
func GetOneUser(id int) (*entities.User, error) {
	var user entities.User
	db := config.GetDB()
	result := db.First(&user, id)
	if result.Error != nil {
		log.Println("Erro ao buscar usuário:", result.Error)
		return nil, result.Error
	}
	return &user, nil
}
func DeleteUser(id int) error {
	db := config.GetDB()
	var user entities.User
	result := db.First(&user, id)
	if result.Error != nil {
		log.Println("Erro ao buscar usuário:", result.Error)
		return result.Error
	}
	db.Delete(&user)
	if result.Error != nil {
		log.Println("Erro ao deletar usuário:", result.Error)
		return result.Error
	}
	log.Println("Usuário deletado com sucesso")
	return nil
}

// Cria o usuário no banco de dados
func CreateUser(firstName, lastName, email string) (*entities.User, error) {
	db := config.GetDB()

	if err := checkUserAndEmail(firstName, lastName, email); err != nil {
		log.Println("Erro de validação:", err)
		return nil, err
	}

	newUser := &entities.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	if err := db.Create(newUser).Error; err != nil {
		log.Println("Erro ao criar usuário:", err)
		return nil, err
	}

	return newUser, nil
}

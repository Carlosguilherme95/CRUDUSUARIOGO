package config

import (
	"crudgo/entities" // importa o pacote onde estão suas entidades
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB inicializa e configura o banco de dados
func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("CRUD.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados:", err)
	}

	// Automigração para manter o schema atualizado
	err = db.AutoMigrate(&entities.User{})
	if err != nil {
		log.Fatal("Erro ao fazer auto migrate:", err)
	}
}

// GetDB retorna a instância do banco
func GetDB() *gorm.DB {
	return db
}

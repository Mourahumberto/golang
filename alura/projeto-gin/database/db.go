package database

import (
	"log"
	"projeto-gin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=54322 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com o Banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}

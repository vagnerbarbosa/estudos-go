package database

import (
	"api-go-gin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=postgres password=1234 dbname=api_go_gin port=5433 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados: ", err)
	}
	DB.AutoMigrate(&models.Aluno{})
}

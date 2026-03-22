package main

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	database.ConectaComBancoDeDados()
	database.DB.AutoMigrate(&models.Personalidade{})
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}

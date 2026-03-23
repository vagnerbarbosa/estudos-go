package main

import (
	"api-go-gin/models"
	"api-go-gin/routes"
)

func main() {

	models.Alunos = []models.Aluno{
		{ID: 1, Nome: "João", CPF: "123.456.789-00", RG: "MG-12.345.678"},
		{ID: 2, Nome: "Maria", CPF: "987.654.321-00", RG: "SP-87.654.321"},
	}

	routes.HandlerRequests()
}

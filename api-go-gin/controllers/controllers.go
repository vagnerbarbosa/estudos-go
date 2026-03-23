package controllers

import (
	"api-go-gin/database"
	"api-go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(http.StatusOK, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Param("nome")
	c.JSON(http.StatusOK, gin.H{
		"API Diz": "Olá, " + nome + "! Seja bem-vindo(a) à API Go com Gin!",
	})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

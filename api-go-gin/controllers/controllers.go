package controllers

import (
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

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   1,
		"nome": "João",
	})
}

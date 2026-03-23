package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   1,
		"nome": "João",
	})
}

func main() {

	r := gin.Default()

	r.GET("/alunos", ExibeTodosAlunos)

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

package routes

import (
	"api-go-gin/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func HandlerRequests() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeTodosAlunos)

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

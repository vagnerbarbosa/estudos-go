package routes

import (
	"api-go-gin/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func HandlerRequests() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.ExibeAlunoPorID)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:CPF", controllers.BuscaAlunoPorCPF)

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

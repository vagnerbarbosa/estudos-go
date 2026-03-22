package routes

import (
	"api-go-rest/controllers"
	"fmt"
	"log"
	"net/http"
)

func HandleRequest() {
	// 1. Substituímos o mux.NewRouter() pelo roteador nativo do Go
	//Poderia usar o DefaultServeMux mas por uma questão de O DefaultServeMux é uma variável global. Se alguma biblioteca de terceiros que você importar decidir registrar uma rota nela maliciosamente.
	r := http.NewServeMux()

	// 2. Definimos as rotas.
	// Note que agora incluímos o método (GET) no início da string.
	r.HandleFunc("GET /", controllers.Home)
	r.HandleFunc("GET /api/personalidades", controllers.TodasPersonalidades)

	r.HandleFunc("GET /api/personalidades/{id}", controllers.RetornaUmaPersonalidade)
	r.HandleFunc("POST /api/personalidades", controllers.CriaUmaPersonalidade)
	r.HandleFunc("PUT /api/personalidades/{id}", controllers.AtualizaUmaPersonalidade)
	r.HandleFunc("DELETE /api/personalidades/{id}", controllers.DeletaUmaPersonalidade)

	fmt.Println("Iniciando o servidor Rest com Go na porta :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

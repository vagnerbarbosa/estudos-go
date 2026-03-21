package main

import (
	"net/http"

	"estudos-go/loja/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}

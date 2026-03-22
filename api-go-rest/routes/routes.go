package routes

import (
	"api-go-rest/controllers"
	"log"
	"net/http"
)

func HandlerRequests() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

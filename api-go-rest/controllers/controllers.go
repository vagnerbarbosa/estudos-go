package controllers

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	result := database.DB.Find(&personalidades)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var personalidade models.Personalidade
	result := database.DB.First(&personalidade, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Personalidade not found", http.StatusNotFound)
		} else {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(personalidade)
}

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

func CriaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	err := json.NewDecoder(r.Body).Decode(&personalidade)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := database.DB.Create(&personalidade)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(personalidade)
}

func AtualizaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var personalidade models.Personalidade
	err := json.NewDecoder(r.Body).Decode(&personalidade)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := database.DB.Model(&models.Personalidade{}).Where("id = ?", id).Updates(personalidade)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	} else if result.RowsAffected == 0 {
		http.Error(w, "Personalidade not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(personalidade)
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	result := database.DB.Delete(&models.Personalidade{}, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "Personalidade not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

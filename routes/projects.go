package routes

import (
	"encoding/json"
	"net/http"

	"github.com/LUISEDOCCOR/api-devprojects/database"
	"github.com/LUISEDOCCOR/api-devprojects/models"
	"github.com/LUISEDOCCOR/api-devprojects/utils"
	"github.com/gorilla/mux"
)

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	var Projects []models.Projects
	database.DB.Find(&Projects)
	json.NewEncoder(w).Encode(Projects)
}

func GetProjectsByCategory(w http.ResponseWriter, r *http.Request) {
	var Projects []models.Projects
	var Category models.Categories
	params := mux.Vars(r)
	IDCategory := params["id"]

	var err error

	err = database.DB.Where("id = ?", IDCategory).First(&Category).Error

	if err != nil {
		response := utils.CreateResponse("error", "There was an error on the server")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	if Category.ID == 0 {
		response := utils.CreateResponse("error", "The desired category was not found")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = database.DB.Where("id_category = ?", Category.ID).Find(&Projects).Error

	if err != nil {
		response := utils.CreateResponse("error", "There was an error on the server")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Projects)
}

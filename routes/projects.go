package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/LUISEDOCCOR/api-devprojects/database"
	"github.com/LUISEDOCCOR/api-devprojects/models"
	"github.com/LUISEDOCCOR/api-devprojects/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
func CreateProject(w http.ResponseWriter, r *http.Request) {

	var err error
	_ = godotenv.Load()

	params := mux.Vars(r)
	app_key := params["app_key"]

	if app_key != os.Getenv("APP_KEY") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var project models.Projects
	var category models.Categories

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&project)

	if err != nil {
		response := utils.CreateResponse("error", "Request error")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if project.Title == "" || project.Content == "" {
		response := utils.CreateResponse("error", "Fill in all fields")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	_ = database.DB.Where("id = ?", project.Id_category).First(&category)

	if category.ID == 0 {
		response := utils.CreateResponse("error", "The category not exists")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = database.DB.Create(&project).Error

	if err != nil {
		response := utils.CreateResponse("error", "Server Error")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := utils.CreateResponse("success", "Successfully created")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

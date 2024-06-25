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

func CreateCategory(w http.ResponseWriter, r *http.Request) {

	_ = godotenv.Load()
	var err error

	params := mux.Vars(r)
	app_key := params["app_key"]

	if app_key != os.Getenv("APP_KEY") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var category models.Categories
	var existingCategory models.Categories

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&category)

	if err != nil {
		response := utils.CreateResponse("error", "Request error")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	_ = database.DB.Where("name = ?", category.Name).First(&existingCategory).Error

	if existingCategory.ID != 0 {
		response := utils.CreateResponse("error", "The category already exists")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if category.Name == "" {
		response := utils.CreateResponse("error", "The name field is necessary")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = database.DB.Create(&category).Error

	if err != nil {
		response := utils.CreateResponse("error", "Request error")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := utils.CreateResponse("success", "Successfully created")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

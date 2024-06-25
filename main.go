package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/LUISEDOCCOR/api-devprojects/database"
	"github.com/LUISEDOCCOR/api-devprojects/models"
	"github.com/LUISEDOCCOR/api-devprojects/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	database.Conn()
	database.DB.AutoMigrate(models.Categories{})
	database.DB.AutoMigrate(models.Projects{})

	// OK
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		var response = map[string]string{
			"status": "ok",
		}
		json.NewEncoder(w).Encode(response)

	})
	//Projects
	r.HandleFunc("/projects", routes.GetAllProjects).Methods("GET")
	//Projects by category
	r.HandleFunc("/projects/{id}", routes.GetProjectsByCategory).Methods("GET")
	//Create category
	r.HandleFunc("/category/{app_key}", routes.CreateCategory).Methods("POST")

	fmt.Println("The server is 🚀")
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		log.Fatal(err)
	}
}

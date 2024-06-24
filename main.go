package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/LUISEDOCCOR/api-devprojects/database"
	"github.com/LUISEDOCCOR/api-devprojects/models"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	database.Conn()
	database.DB.AutoMigrate(models.Categories{})
	database.DB.AutoMigrate(models.Projects{})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		var response = map[string]string{
			"status": "ok",
		}
		json.NewEncoder(w).Encode(response)

	})

	fmt.Println("The server is 🚀")
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	persistence "github.com/gegen07/cartola-university/infrastructure/persistance"
	"github.com/gegen07/cartola-university/interfaces"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env")
	}
}

func main() {
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbpass := os.Getenv("DB_PASSWORD")
	dbuser := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	dbdriver := os.Getenv("DB_DRIVER")

	services, err := persistence.NewRepositories(dbdriver, dbuser, dbpass, dbport, dbhost, dbname)

	if err != nil {
		panic(err)
	}

	_ = services.Migrate()

	formations := interfaces.NewFormationHandler(services.Formation)

	router := mux.NewRouter()

	//formation routes
	router.HandleFunc("/formation", formations.Insert).Methods("POST")
	router.HandleFunc("/formation", formations.GetAllFormations).Methods("GET")
	router.HandleFunc("/formation/{id}", formations.GetFormationById).Methods("GET")
	router.HandleFunc("/formation/{id}", formations.Update).Methods("PUT")
	router.HandleFunc("/formation/{id}", formations.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(router)))
}

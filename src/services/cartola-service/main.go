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
	scouts := interfaces.NewScoutHandler(services.Scout)

	router := mux.NewRouter()

	//formation routes
	router.Handle("/formation", interfaces.RootHandler(formations.Insert)).Methods(http.MethodPost)
	router.Handle("/formation", interfaces.RootHandler(formations.GetAllFormations)).Methods(http.MethodGet)
	router.Handle("/formation/{id}", interfaces.RootHandler(formations.GetFormationById)).Methods(http.MethodGet)
	router.Handle("/formation/{id}", interfaces.RootHandler(formations.Update)).Methods(http.MethodPut)
	router.Handle("/formation/{id}", interfaces.RootHandler(formations.Delete)).Methods(http.MethodDelete)

	//scout routes
	router.Handle("/scout", interfaces.RootHandler(scouts.Insert)).Methods(http.MethodPost)
	router.Handle("/scout", interfaces.RootHandler(scouts.GetAllScouts)).Methods(http.MethodGet)
	router.Handle("/scout/{id}", interfaces.RootHandler(scouts.GetScoutByID)).Methods(http.MethodGet)
	router.Handle("/scout/{id}", interfaces.RootHandler(scouts.Update)).Methods(http.MethodPut)
	router.Handle("/scout/{id}", interfaces.RootHandler(scouts.Delete)).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(router)))
}

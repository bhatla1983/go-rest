package main

import (
	"go-rest/controllers"
	"go-rest/database"
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	// database_user := os.Args[1]
	// database_password := os.Args[2]
	// fmt.Println("DB User: ", database_user)
	// fmt.Println("DB Password: ", database_password)
	dbhost := os.Getenv("DB_SERVER_NAME")
	log.Println("DB server name from Env : ", dbhost)

	initDB()
	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)
	initialiseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initialiseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllPerson).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletPersonByID).Methods("DELETE")
}

func initDB() {
	config, err := LoadAppConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	conn :=
		database.Config{
			ServerName: config.DBServerName,
			User:       config.DBUser,
			Password:   config.DBPassword,
			DB:         config.DBName,
		}
	connectionString := database.GetConnectionString(conn)
	err = database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
}

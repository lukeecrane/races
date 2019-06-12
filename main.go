package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"webstuff/go-contacts/app"
	"webstuff/go-contacts/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts

	//Create and saving a new group login for race registeration
	router.HandleFunc("/api/group/new", controllers.CreateGroup).Methods("POST")
	router.HandleFunc("/api/group/login", controllers.AuthenticateGroup).Methods("POST")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, handlers.CORS()(router)) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}

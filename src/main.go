package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	controllers "./controllers"
)

func main() {
	// Specifies application port.
	port := ":3005"
	router := mux.NewRouter()

	// Specifies routes.
	router.HandleFunc("/videos", controllers.GetVideos).Methods("GET")
	// Allows Cross Origin Resource Sharing.

	// Specifies HTTP methods available.
	allowedMethods := handlers.AllowedMethods([]string{"GET"})

	if err := http.ListenAndServe(port,
		handlers.CORS(allowedMethods)(router)); err != nil {
		log.Fatal(err)
	}

}

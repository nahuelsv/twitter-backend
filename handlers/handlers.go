package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nahuelsv/twitter-backend/middlewares"
	"github.com/nahuelsv/twitter-backend/routers"
	"github.com/rs/cors"
)

/*managers is a method that pass a router object to a cors handler to serv our server*/
func Managers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

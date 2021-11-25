package routes

import (
	// "fmt"
	// "net/http"

	"github.com/go-chat/server/controllers"
	"github.com/go-chat/server/middlewares"
	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Index).Methods("GET")	
	r.HandleFunc("/login", middlewares.CheckJwt(controllers.Login)).Methods("GET", "POST")
	r.HandleFunc("/register", middlewares.CheckJwt(controllers.Register)).Methods("GET", "POST")
	return r
}
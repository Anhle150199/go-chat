package routes

import (
	// "fmt"
	// "net/http"

	"github.com/go-chat/server/controllers"
	// "github.com/go-chat/server/middlewares"
	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Index).Methods("GET")

	authRouter := r.PathPrefix("/").Subrouter()
	authRouter.HandleFunc("/login", controllers.Login).Methods("GET", "POST")
	authRouter.HandleFunc("/register", controllers.Register).Methods("GET", "POST")
	authRouter.HandleFunc("/logout", controllers.Logout).Methods("POST")
	return r
}

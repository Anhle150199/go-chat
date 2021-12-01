package routes

import (
	"net/http"

	"github.com/go-chat/server/controllers"
	"github.com/go-chat/server/middlewares"
	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	authRouter := r.PathPrefix("/").Subrouter()
	authRouter.HandleFunc("/login", controllers.Login).Methods("GET", "POST")
	authRouter.HandleFunc("/register", controllers.Register).Methods("GET", "POST")
	authRouter.HandleFunc("/logout", controllers.Logout).Methods("POST")

	chatRouter := r.PathPrefix("/").Subrouter()
	chatRouter.Use(middlewares.CheckJwt)
	chatRouter.HandleFunc("/chat", controllers.Index).Methods("GET")
	chatRouter.HandleFunc("/edit-name", controllers.EditName).Methods("POST")

	return r
}

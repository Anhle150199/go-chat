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

	authRouter := r.PathPrefix("/user").Subrouter()
	authRouter.HandleFunc("/login", controllers.Login).Methods("GET", "POST")
	authRouter.HandleFunc("/register", controllers.Register).Methods("GET", "POST")
	authRouter.HandleFunc("/logout", controllers.Logout).Methods("POST")

	chatRouter := r.PathPrefix("/chat").Subrouter()
	chatRouter.Use(middlewares.CheckJwt)
	chatRouter.HandleFunc("/chat", controllers.Index).Methods("GET")
	chatRouter.HandleFunc("/edit-name", controllers.EditName).Methods("POST")
	chatRouter.HandleFunc("/sent-message", controllers.SendMessage).Methods("POST")
	chatRouter.HandleFunc("/edit-message", controllers.EditMessage).Methods("POST")
	chatRouter.HandleFunc("/delete-message", controllers.DeleteMessage).Methods("POST")

	return r
}

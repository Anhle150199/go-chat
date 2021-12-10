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

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/chat/", http.StatusMovedPermanently)
	})

	authRouter := r.PathPrefix("/user").Subrouter()
	authRouter.HandleFunc("/login", controllers.Login).Methods("GET")
	authRouter.HandleFunc("/login", controllers.Login).Methods("POST")
	authRouter.HandleFunc("/register", controllers.Register).Methods("GET")
	authRouter.HandleFunc("/register", controllers.Register).Methods("POST")
	authRouter.HandleFunc("/logout", controllers.Logout).Methods("POST")

	chatRouter := r.PathPrefix("/chat").Subrouter()
	chatRouter.Use(middlewares.CheckJwt)
	chatRouter.HandleFunc("/", controllers.Index).Methods("GET")
	chatRouter.HandleFunc("/edit-name", controllers.EditName).Methods("POST")
	chatRouter.HandleFunc("/sent-message", controllers.SendMessage).Methods("POST")
	chatRouter.HandleFunc("/edit-message", controllers.EditMessage).Methods("POST")
	chatRouter.HandleFunc("/delete-message", controllers.DeleteMessage).Methods("POST")
	chatRouter.HandleFunc("/upload-media", controllers.UploadFile).Methods("POST")
	chatRouter.HandleFunc("/sent-stamp", controllers.SentStamp).Methods("POST")

	return r
}

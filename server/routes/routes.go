package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/go-chat/server/controllers"

)
func Setup() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintln(w,"GO CHAT")
	}).Methods("GET")	
	r.HandleFunc("/login", controllers.Login).Methods("GET")
	return r
}
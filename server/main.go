package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chat/server/routes"
	"github.com/go-chat/server/database"

	_ "github.com/gorilla/mux"
)
func main()  {
	fmt.Println("Starting server 1.0.0")

	database.ConnectDB()
	

	r := routes.Setup()
	port:= ":80"
	log.Println(fmt.Sprintf("server is running at: http://127.0.0.1%v", port))
	log.Fatalln(http.ListenAndServe(port, r))

}